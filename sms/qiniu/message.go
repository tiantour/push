package qiniu

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/duke-git/lancet/v2/datetime"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/google/go-querystring/query"
)

// Message message
type Message struct {
	Phone    []string `json:"phone,omitempty"`    // 号码
	Template string   `json:"template,omitempty"` // 模板
	// Signature string            `json:"signature,omitempty"` // 模板
	Body map[string]string `json:"body,omitempty"` // 参数
	Date string            `json:"date,omitempty"` // 日期
	Page int               `json:"page,omitempty"` // 页码
	Size int               `json:"size,omitempty"` // 数量
}

// NewMessage new message
func NewMessage() *Message {
	return &Message{}
}

// Query query message
func (m *Message) Query(args *Message) ([]*QueryMessage, error) {
	date, err := datetime.FormatStrToTime(args.Date, "yyyy-mm-dd")
	if err != nil {
		return nil, err
	}

	requst := &QueryRequest{
		Mobile:   args.Phone[0],
		Start:    fmt.Sprintf("%d", datetime.BeginOfDay(date).Unix()),
		End:      fmt.Sprintf("%d", datetime.EndOfDay(date).Unix()),
		Page:     args.Page + 1,
		PageSize: args.Size,
	}

	signURL, err := query.Values(requst)
	if err != nil {
		return nil, err
	}

	body := []byte{}

	header := http.Header{}
	header.Add("Content-Length", fmt.Sprintf("%d", len(body)))
	header.Add("Content-Type", "application/json")

	data := netutil.HttpRequest{
		RawURL:  "https://sms.qiniuapi.com/v1/message?" + signURL.Encode(),
		Method:  "GET",
		Headers: header,
	}

	sign, err := NewSMS().Sign(&data)
	if err != nil {
		return nil, err
	}

	data.Headers.Add("Authorization", fmt.Sprintf("Qiniu %s:%s", SecretKey, sign))

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&data)
	if err != nil {
		return nil, err
	}

	result := QueryResponse{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	return result.Items, err
}

// Send send message
func (m *Message) Send(args *Message) (*SendResponse, error) {
	request := SendRequest{
		Mobiles:    args.Phone,
		Parameters: args.Body,
		TemplateID: args.Template,
	}

	body, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	header := http.Header{}
	header.Add("Content-Length", fmt.Sprintf("%d", len(body)))
	header.Add("Content-Type", "application/json")

	data := netutil.HttpRequest{
		RawURL:  "https://sms.qiniuapi.com/v1/message",
		Method:  "POST",
		Headers: header,
		Body:    body,
	}

	sign, err := NewSMS().Sign(&data)
	if err != nil {
		return nil, err
	}

	// urlsafe
	sign = strings.ReplaceAll(sign, "+", "-")
	sign = strings.ReplaceAll(sign, "/", "_")

	// AccessKey
	data.Headers.Add("Authorization", fmt.Sprintf("Qiniu %s:%s", AccessKey, sign))

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&data)
	if err != nil {
		return nil, err
	}

	result := SendResponse{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.Error != "" {
		return nil, errors.New(result.Message)
	}
	return &result, err

}
