package qiniu

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/tiantour/fetch"
	"github.com/tiantour/tempo"
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
	start, err := tempo.NewString().Unix(fmt.Sprintf("%s 00:00:00", args.Date))
	if err != nil {
		return nil, err
	}
	end, err := tempo.NewString().Unix(fmt.Sprintf("%s 23:59:59", args.Date))
	if err != nil {
		return nil, err
	}
	requst := &QueryRequest{
		Mobile:   args.Phone[0],
		Start:    fmt.Sprintf("%d", start),
		End:      fmt.Sprintf("%d", end),
		Page:     args.Page + 1,
		PageSize: args.Size,
	}
	signURL, err := query.Values(requst)
	if err != nil {
		return nil, err
	}

	body := []byte{}
	header := http.Header{
		"Content-Length": []string{fmt.Sprintf("%d", len(body))},
		"Content-Type":   []string{"application/json"},
	}
	data := fetch.Request{
		Method: "GET",
		URL:    "https://sms.qiniuapi.com/v1/message?" + signURL.Encode(),
		Header: header,
	}
	sign, err := NewSMS().Sign(&data)
	if err != nil {
		return nil, err
	}
	token := fmt.Sprintf("Qiniu %s:%s", SecretKey, sign)
	header.Add("Authorization", token)
	body, err = fetch.Cmd(&data)
	if err != nil {
		return nil, err
	}

	result := QueryResponse{}
	err = json.Unmarshal(body, &result)
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

	header := http.Header{
		"Content-Length": []string{fmt.Sprintf("%d", len(body))},
		"Content-Type":   []string{"application/json"},
	}

	data := fetch.Request{
		Method: "POST",
		URL:    "https://sms.qiniuapi.com/v1/message",
		Body:   body,
		Header: header,
	}

	sign, err := NewSMS().Sign(&data)
	if err != nil {
		return nil, err
	}

	// urlsafe
	sign = strings.ReplaceAll(sign, "+", "-")
	sign = strings.ReplaceAll(sign, "/", "_")

	// AccessKey
	token := fmt.Sprintf("Qiniu %s:%s", AccessKey, sign)
	header.Add("Authorization", token)

	body, err = fetch.Cmd(&data)
	if err != nil {
		return nil, err
	}

	result := SendResponse{}
	err = json.Unmarshal(body, &result)
	return &result, err

}
