package sms

import (
	"encoding/json"
	"time"
)

// Send send
type Send struct {
	Extend          string `json:"extend" url:"extend,omitempty"`                         // 否 拓展 会员id
	SmsType         string `json:"sms_type" url:"sms_type,omitempty"`                     // 是 类型 normal
	SmsFreeSignName string `json:"sms_free_sign_name" url:"sms_free_sign_name,omitempty"` // 是 短信签名
	SmsParam        string `json:"sms_param" url:"sms_param,omitempty"`                   // 否 模板变量
	RecNum          string `json:"rec_num" url:"rec_num,omitempty"`                       // 是 接收号码，英文逗号分割，最大200
	SmsTemplateCode string `json:"sms_template_code" url:"sms_template_code,omitempty"`   // 是 模板ID
	Request
}

// NewSend new send
func NewSend() *Send {
	return &Send{}
}

// Do do
func (s *Send) Do(phone, template string, content map[string]string) ([]byte, error) {
	body, err := json.Marshal(content)
	if err != nil {
		return []byte{}, err
	}
	data := Send{
		SmsType:         "normal",
		SmsFreeSignName: Sign,
		SmsParam:        string(body),
		RecNum:          phone,
		SmsTemplateCode: template,
		Request: Request{
			Method:     "alibaba.aliqin.fc.sms.num.send",
			AppKey:     AppKey,
			SignMethod: "md5",
			Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
			Format:     "json",
			V:          "2.0",
		},
	}
	return NewSMS().Body(data)
}
