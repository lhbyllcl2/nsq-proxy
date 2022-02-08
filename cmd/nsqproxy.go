package main

//go:generate echo "statik -src=../web/vue-admin/dist/ -dest=../internal -f"
//go:generate statik -src=../web/vue-admin/dist/ -dest=../internal -f

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arl/statsviz"
	"nsq-proxy/config"
	"nsq-proxy/internal/backup"
	"nsq-proxy/internal/httper"
	"nsq-proxy/internal/model"
	"nsq-proxy/internal/module/logger"
	"nsq-proxy/internal/module/tool"
	"nsq-proxy/internal/proxy"
	"nsq-proxy/internal/publish"
)

// 主函数
func main() {
	//初始化系统配置
	config.NewSystemConfig()
	//连接数据库
	model.NewDB(config.SystemConfig.DbHost, config.SystemConfig.DbPort, config.SystemConfig.DbUsername, config.SystemConfig.DbPassword, config.SystemConfig.DbName)
	//连接生产者
	publish.InitProducer(config.SystemConfig.NsqTcpAddress)
	//创建一个proxy实例
	p := proxy.NewProxy()
	//异常捕获
	defer func() {
		tool.PanicHandlerForLog()
		logger.Fatalf("nsqProxy will exit")
		os.Exit(2)
	}()
	//开启HTTP
	httper.NewHttper(config.SystemConfig.HttpAddr).Run()
	//开启监控
	_ = statsviz.RegisterDefault()
	//灾备
	backup.Backup(config.SystemConfig.MasterAddr)
	//启动一个proxy实例
	logger.Infof("nsqProxy is starting")
	p.Run()
	fmt.Println("service start...")
	//监听信号
	listenSignal(p)
	logger.Infof("nsqProxy end success")
}

// 监听信号
func listenSignal(p *proxy.Proxy) {
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTRAP)
	for {
		sig := <-sigChannel
		logger.Infof("nsqProxy receive signal: %s", sig.String())
		if sig == syscall.SIGTRAP {
			continue
		}
		logger.Infof("nsqProxy is closing consumes...")
		p.SetExitFlag()
		p.Stop()
		publish.Producer.Stop()
		//等待10秒
		logger.Infof("nsqProxy will be closed master process ten seconds later.")
		time.Sleep(10)
		//time.Sleep(10 * time.Second)
		break
	}
}
