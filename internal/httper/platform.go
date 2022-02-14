package httper

import (
	"net/http"
	"strconv"

	"nsq-proxy/config"
	"nsq-proxy/internal/model"
	"nsq-proxy/internal/module/tool"
)

type platform struct {
}

func NewPlatform() *platform {
	return &platform{}
}
func (slf *platform) Create(w http.ResponseWriter, r *http.Request) {
	pf := &model.Platform{}
	pf.Name = r.FormValue("name")
	pf.Remark = r.FormValue("remark")
	status, err := strconv.Atoi(r.FormValue("status"))
	if err != nil {
		status = 1
	}
	pf.Status = status
	firstId := pf.FirstByName(pf.Name)
	if firstId > 0 {
		Failed(w, 408, "名称已存在")
		return
	}
	id, err := pf.Create()
	if err != nil {
		Failed(w, HttpCodeBadRequest, "create failed. err: "+err.Error())
		return
	}
	if id <= 0 {
		Failed(w, HttpCodeBadRequest, "id is zero")
		return
	}
	_, err = pf.Update(map[string]interface{}{
		"app_id":     tool.AppIdEncode(id, config.HdSalt),
		"app_secret": tool.AppSecretEncode(id, config.HdSalt),
	})
	if err != nil {
		return
	}
	Success(w, pf)
}

func (slf *platform) Page(w http.ResponseWriter, r *http.Request) {
	ms := &model.Platform{}
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	AppId := r.FormValue("AppId")
	status := r.FormValue("status")
	pageResult, err := ms.Page(page, AppId, status)
	if err != nil {
		Failed(w, HttpCodeInternalServerError, "please try again. err: "+err.Error())
		return
	}
	Success(w, pageResult)
}
func (slf *platform) Delete(w http.ResponseWriter, r *http.Request) {
	pf := &model.Platform{}
	var err error
	pf.Id, err = strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil || pf.Id <= 0 {
		Failed(w, HttpCodeBadRequest, "param error")
		return
	}
	_, err = pf.Delete()
	if err != nil {
		Failed(w, HttpCodeBadRequest, "delete failed. err: "+err.Error())
		return
	}
	Success(w, "ok")
}
