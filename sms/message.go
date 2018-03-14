package sms

import (
	"strings"
	"time"
)

type (
	// Send send
	Send struct {
		Extend          string `json:"extend" url:"extend,omitempty"`                         // 否 拓展 会员id
		SmsType         string `json:"sms_type" url:"sms_type,omitempty"`                     // 是 类型 normal
		SmsFreeSignName string `json:"sms_free_sign_name" url:"sms_free_sign_name,omitempty"` // 是 短信签名
		SmsParam        string `json:"sms_param" url:"sms_param,omitempty"`                   // 否 模板变量
		RecNum          string `json:"rec_num" url:"rec_num,omitempty"`                       // 是 接收号码，英文逗号分割，最大200
		SmsTemplateCode string `json:"sms_template_code" url:"sms_template_code,omitempty"`   // 是 模板ID
		Request
	}

	// Query query
	Query struct {
		BizID       string `json:"biz_id" url:"biz_id,omitempty"`             // 否 流水号
		RecNum      string `json:"rec_num" url:"rec_num,omitempty"`           // 是 手机号码
		QueryDate   string `json:"query_date" url:"query_date,omitempty"`     // 是 日期 20170525
		CurrentPage int    `json:"current_page" url:"current_page,omitempty"` // 是 页码
		PageSize    int    `json:"page_size" url:"page_size,omitempty"`       // 是 数量
		Request
	}
)

// Message message
type Message struct {
	Phone    string `json:"phone"`    // 号码
	Template string `json:"template"` // 模板
	Body     []byte `json:"data"`     // 内容
	Date     string `json:"date"`     // 日期
	Page     int    `json:"page"`     // 页码
	Size     int    `json:"size"`     // 数量
}

// NewMessage new message
func NewMessage() *Message {
	return &Message{}
}

// Send send
func (m *Message) Send(args *Message) (*Response, error) {
	data := &Send{
		SmsType:         "normal",
		SmsFreeSignName: Sign,
		SmsParam:        string(args.Body),
		RecNum:          args.Phone,
		SmsTemplateCode: args.Template,
		Request: Request{
			Method:     "alibaba.aliqin.fc.sms.num.send",
			AppKey:     AppKey,
			SignMethod: "md5",
			Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
			Format:     "json",
			V:          "2.0",
		},
	}
	return NewSMS().do(data)
}

// Query query message
func (m *Message) Query(args *Message) (*Response, error) {
	data := &Query{
		RecNum:      args.Phone,
		QueryDate:   strings.Replace(args.Date, "-", "", -1),
		CurrentPage: args.Page,
		PageSize:    args.Size,
		Request: Request{
			Method:     "alibaba.aliqin.fc.sms.num.query",
			AppKey:     AppKey,
			SignMethod: "md5",
			Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
			Format:     "json",
			V:          "2.0",
		},
	}
	return NewSMS().do(data)
}
