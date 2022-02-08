package httper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
	"nsq-proxy/internal/model"
	"nsq-proxy/internal/module/logger"
	"nsq-proxy/internal/module/tool"
	"nsq-proxy/internal/publish"
)

type publisher struct {
	secret  string
	publish *nsq.Producer
}

func NewPublisher() *publisher {
	return &publisher{
		secret:  "",
		publish: publish.Producer,
	}
}

type RequestParams struct {
	AppId     string `json:"app_id"`
	Topic     string `json:"topic"`
	Delay     int64  `json:"delay"`
	Signature string `json:"signature"`
	Timestamp string `json:"timestamp"`
	model.Message
}

func (slf *publisher) Create(w http.ResponseWriter, r *http.Request) {
	params := RequestParams{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		logger.Infof("[error] decoding body: %s", err)
		Failed(w, HttpCodeBadRequest, "Error parsing request body")
		return
	}
	if err := tool.Validate(params); err != nil {
		logger.Infof("[error] decoding body: %s", err)
		Failed(w, HttpCodeBadRequest, err.Error())
		return
	}
	signerMap := map[string]string{
		"app_id":    params.AppId,
		"timestamp": params.Timestamp,
		"topic":     params.Topic,
		"delay":     strconv.Itoa(int(params.Delay)),
		"url":       params.URL,
		"method":    params.Method,
		"body":      params.Body,
	}
	sign := tool.Sign(signerMap, "")
	fmt.Println(sign)
	id := tool.GenerateUniqueId(2)
	messageID := fmt.Sprintf("%s", id)
	msg := &model.Message{
		ID:     messageID,
		URL:    params.URL,
		Method: params.Method,
		Body:   params.Body,
	}
	msgContent, err := json.Marshal(msg)
	if err != nil {
		logger.Infof("[error] marshalling message: %s", err)
		Failed(w, HttpCodeBadRequest, "Error parsing request body")
		return
	}
	if params.Delay > 0 {
		seconds := time.Duration(params.Delay)
		if err = slf.publish.DeferredPublish(params.Topic, seconds*time.Second, msgContent); err != nil {
			Failed(w, HttpCodeBadRequest, err.Error())
		}
		mod := &model.NsqproxyMessage{
			MessageID: messageID,
			Topic:     params.Topic,
			URL:       params.URL,
			Method:    params.Method,
			Argument:  params.Body,
			Delay:     params.Delay,
		}
		_ = mod.Create()
	} else {
		if err = slf.publish.Publish(params.Topic, msgContent); err != nil {
			Failed(w, HttpCodeBadRequest, err.Error())
		}
	}
	Success(w, "ok")
}
