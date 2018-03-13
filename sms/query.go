package sms

import (
	"strings"
	"time"
)

type (
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

// NewQuery new query
func NewQuery() *Query {
	return &Query{}
}

// Do do
func (q *Query) Do(phone, date string, page, size int) ([]byte, error) {
	data := Query{
		RecNum:      phone,
		QueryDate:   strings.Replace(date, "-", "", -1),
		CurrentPage: page,
		PageSize:    size,
		Request: Request{
			Method:     "alibaba.aliqin.fc.sms.num.query",
			AppKey:     AppKey,
			SignMethod: "md5",
			Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
			Format:     "json",
			V:          "2.0",
		},
	}
	return NewSMS().Body(data)
}
