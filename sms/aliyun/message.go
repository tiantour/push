package aliyun

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/duke-git/lancet/v2/random"
	"github.com/google/go-querystring/query"
)

// Message message
type Message struct {
	Phone    []string          `json:"phone,omitempty"`    // 号码
	Template string            `json:"template,omitempty"` // 模板
	Body     map[string]string `json:"body,omitempty"`     // 参数
	Date     string            `json:"date,omitempty"`     // 日期
	Page     int               `json:"page,omitempty"`     // 页码
	Size     int               `json:"size,omitempty"`     // 数量
}

// NewMessage new message
func NewMessage() *Message {
	return &Message{}
}

// Query query message
func (m *Message) Query(args *Message) ([]*QueryResponseItem, error) {
	data := &QueryRequest{
		AccessKeyID: AccessKeyID,
		Action:      "QuerySendDetails",
		CurrentPage: args.Page + 1,
		PageSize:    args.Size,
		PhoneNumber: args.Phone[0],
		SendDate:    strings.Replace(args.Date, "-", "", -1),
		Request: Request{
			Format:           "JSON",
			Version:          "2017-05-25",
			Timestamp:        time.Now().In(time.FixedZone("GMT", 0)).Format("2006-01-02T15:04:05Z"),
			SignatureNonce:   random.RandString(32),
			SignatureMethod:  "HMAC-SHA1",
			SignatureVersion: "1.0",
			RegionID:         "cn-hangzhou",
		},
	}

	signURL, err := query.Values(data)
	if err != nil {
		return nil, err
	}

	sign := NewSMS().Sign(signURL)
	signURL.Add("Signature", sign)

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://dysmsapi.aliyuncs.com?%s", signURL.Encode()),
		Method: "POST",
		Headers: http.Header{
			"Content-Type": []string{"application/json"},
		},
	})
	if err != nil {
		return nil, err
	}

	result := QueryResponse{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.Code != "OK" {
		return nil, errors.New(result.Message)
	}
	return result.SmsSendDetailDTOs.SmsSendDetailDTO, nil
}

// Send send
func (m *Message) Send(args *Message) (*Response, error) {
	body, err := json.Marshal(args.Body)
	if err != nil {
		return nil, err
	}

	data := &SendRequest{
		AccessKeyID:   AccessKeyID,
		Action:        "SendSms",
		PhoneNumbers:  strings.Join(args.Phone, ","),
		SignName:      Sign,
		TemplateCode:  args.Template,
		TemplateParam: string(body),
		Request: Request{
			Format:           "JSON",
			Version:          "2017-05-25",
			Timestamp:        time.Now().In(time.FixedZone("GMT", 0)).Format("2006-01-02T15:04:05Z"),
			SignatureNonce:   random.RandString(32),
			SignatureMethod:  "HMAC-SHA1",
			SignatureVersion: "1.0",
			RegionID:         "cn-hangzhou",
		},
	}

	signURL, err := query.Values(data)
	if err != nil {
		return nil, err
	}

	sign := NewSMS().Sign(signURL)
	signURL.Add("Signature", sign)

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://dysmsapi.aliyuncs.com?%s", signURL.Encode()),
		Method: "POST",
		Headers: http.Header{
			"Content-Type": []string{"application/json"},
		},
	})
	if err != nil {
		return nil, err
	}

	result := SendResponse{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.Code != "OK" {
		return nil, errors.New(result.Message)
	}
	return &result.Response, nil
}
