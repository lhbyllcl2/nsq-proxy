package model

import (
	"strings"
	"time"
)

type NsqproxyMessage struct {
	ID        uint   `gorm:"column:id;primary_key" json:"id"`
	MessageID string `json:"message_id"`
	Topic     string `json:"topic"`
	URL       string `json:"url"`
	Method    string `json:"method"`
	Delay     int64  `json:"delay"`
	Argument  string `json:"argument"`
	Type      string `json:"type"`
	Status    int    `json:"status"`
	Response  string `json:"response"`
}
type NsqproxyMessagePage struct {
	NsqproxyMessage
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (NsqproxyMessage) TableName() string {
	return `nsqproxy_message`
}

func (NsqproxyMessage) CreateTable() error {
	sql := `
		CREATE TABLE nsqproxy_message (
		  id bigint(20) unsigned AUTO_INCREMENT,
		  message_id varchar(32) NOT NULL DEFAULT '',
		  topic varchar(32) NOT NULL,
		  url varchar(200) NOT NULL,
		  method varchar(10) NOT NULL DEFAULT '',
          delay int(11) unsigned NOT NULL DEFAULT '0' COMMENT '延时多少秒执行',
		  argument json,
		  status tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '结果标识 0-等待执行 1-执行成功  2-执行失败',
		  response varchar(255),
		  create_at timestamp NULL  DEFAULT current_timestamp(),
		  update_at timestamp NULL  DEFAULT current_timestamp(),
		  PRIMARY KEY (id)
		) engine=InnoDB DEFAULT charset=utf8mb4 COMMENT='消息表';
`
	return db.Exec(sql).Error
}
func (slf *NsqproxyMessage) Create() error {
	result := db.FirstOrCreate(slf, NsqproxyMessage{MessageID: slf.MessageID})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func UpdateMessage(messageId string, isSuccess int, response string) error {
	// 条件更新
	if len(response) > 255 {
		response = response[:255]
	}
	result := db.Model(&NsqproxyMessage{}).Where("message_id = ?", messageId).Updates(map[string]interface{}{
		"response":   response,
		"is_success": isSuccess,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (slf *NsqproxyMessage) Page(page int, topic string, status string, Type string) (PageResult, error) {
	var mList []NsqproxyMessagePage
	whereList := make([]string, 0)
	if topic != "" {
		whereList = append(whereList, " topic = '"+topic+"'")
	}
	if status != "" {
		whereList = append(whereList, " status = "+status)
	}
	if Type != "" {
		whereList = append(whereList, " type = "+Type)
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
