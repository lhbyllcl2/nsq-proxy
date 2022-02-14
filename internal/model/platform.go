package model

import (
	"errors"
	"strings"
)

type Platform struct {
	Id        int64  `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Topic     string `json:"topic"`
	Status    int    `json:"status"`
	Remark    string `json:"remark"`
}

func (Platform) TableName() string {
	return `nsqproxy_platform` //
}
func (slf *Platform) Page(page int, AppId string, status string) (PageResult, error) {
	var mList []Platform
	whereList := make([]string, 0)
	whereList = append(whereList, " is_delete = 0")
	if status != "" {
		whereList = append(whereList, " status = "+status)
	}
	if AppId != "" {
		whereList = append(whereList, " app_id = "+AppId)
	}
	d := db.Table(slf.TableName()).Where(strings.Join(whereList, " AND "))
	//count部分
	var total int64
	result := d.Count(&total)
	if result.Error != nil || result.RowsAffected != 1 {
		total = 0
	}
	//page部分
	if page <= 0 {
		page = 1
	}
	result = d.Offset((page - 1) * 20).Order("id desc").Limit(20).Find(&mList)
	pageRet := PageResult{
		Total:  total,
		Page:   page,
		Result: mList,
	}
	return pageRet, result.Error
}
func (slf *Platform) Create() (int64, error) {
	result := db.Create(slf)
	if result.Error != nil {
		return 0, result.Error
	} else if result.RowsAffected <= 0 {
		return 0, errors.New("RowsAffected is zero")
	} else if slf.Id <= 0 {
		return 0, errors.New("primaryKey is zero")
	}
	return slf.Id, nil
}
func (slf *Platform) Update(updateData map[string]interface{}) (int64, error) {
	if slf.Id <= 0 {
		return 0, errors.New("primaryKey is zero")
	}
	result := db.Model(slf).Updates(updateData)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
func (slf *Platform) FirstByName(name string) int64 {
	db.Where("name = ?", name).First(slf)
	return slf.Id
}
func (slf *Platform) Delete() (int64, error) {
	if slf.Id <= 0 {
		return 0, errors.New("primaryKey is zero")
	}
	result := db.Model(slf).Updates(map[string]interface{}{
		"is_delete": 1,
	})
	return result.RowsAffected, result.Error
}
