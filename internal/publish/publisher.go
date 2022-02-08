package publish

import (
	"time"

	"github.com/nsqio/go-nsq"
	"nsq-proxy/internal/module/logger"
)

var Producer *nsq.Producer

func InitProducer(nsqDString string) {
	connectProducer(nsqDString)
	go func(nsqDString string) {
		ticker := time.NewTicker(time.Second * 10)
		defer ticker.Stop()
		for range ticker.C {
			if Producer == nil {
				logger.Fatalf("nsqd retry the connection")
				connectProducer(nsqDString)
			}
		}
	}(nsqDString)
}
func connectProducer(nsqDString string) {
	prod, err := nsq.NewProducer(nsqDString, nsq.NewConfig())
	if err != nil {
		logger.Infof("[error] starting producer: %s", err)
		return
	}
	Producer = prod
}
