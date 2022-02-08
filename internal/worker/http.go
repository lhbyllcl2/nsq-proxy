package worker

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/nsqio/go-nsq"
	"nsq-proxy/internal/model"
	"nsq-proxy/internal/module/tool"
)

type HTTPWorker struct {
	workerConfig workerConfig
	clientPool   *tool.HttpClientPool
}

func (w *HTTPWorker) new(wc workerConfig) {
	w.workerConfig = wc
	w.clientPool = tool.NewHttpClientPool()
}

// Send 给HTTP发消息
func (w *HTTPWorker) Send(message *nsq.Message) ([]byte, error) {
	//构造HTTP请求
	//values := url.Values{}
	//values.Set("param", string(message.Body))
	var ms model.Message
	err := json.Unmarshal(message.Body, &ms)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(ms.Method, ms.URL, strings.NewReader(ms.Body))
	if err != nil {
		return nil, err
	}
	//含下划线会被nginx抛弃，横线会被转为下划线。
	/*	req.Header.Set("MESSAGE_ID", string(message.ID[:]))
		req.Header.Set("MESSAGE-ID", string(message.ID[:]))*/
	req.Header.Set("CONTENT-TYPE", "application/x-www-form-urlencoded")
	//获取http.Client
	client := w.clientPool.GetClient()
	if client == nil {
		return nil, errors.New("HttpClientPool.GetClient is nil")
	}
	defer w.clientPool.PutClient(client)
	client.Timeout = w.workerConfig.timeoutDial
	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, newWorkerErrorWrite(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, newWorkerErrorRead(err)
	}
	return content, nil
}
