package model

type Platform struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Topic     string `json:"topic"`
	Remark    string `json:"remark"`
}

func (Platform) TableName() string {
	return `nsqproxy_platform` //
}
