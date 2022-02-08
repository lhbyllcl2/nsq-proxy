package httper

import (
	"net/http"
	"strconv"

	"nsq-proxy/internal/model"
)

type message struct {
}

func NewMessage() *message {
	return &message{}
}
func (slf *message) Page(w http.ResponseWriter, r *http.Request) {
	ms := &model.NsqproxyMessage{}
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	topic := r.FormValue("topic")
	status := r.FormValue("status")
	msgType := r.FormValue("msgType")
	pageResult, err := ms.Page(page, topic, status, msgType)
	if err != nil {
		Failed(w, HttpCodeInternalServerError, "please try again. err: "+err.Error())
		return
	}
	Success(w, pageResult)
}
