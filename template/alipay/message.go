package alipay

import (
	"errors"
	"fmt"

	"github.com/duke-git/lancet/v2/datetime"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/google/go-querystring/query"
)

// Message message
type Message struct{}

// NewMessage new message
func NewMessage() *Message {
	return &Message{}
}

// MI qrcode
func (m *Message) MI(content string) (*Response, error) {
	args := &Request{
		AppID:      AppID,
		Method:     "alipay_open_app_mini_templatemessage_send_response",
		Format:     "JSON",
		Charset:    "utf-8",
		SignType:   "RSA2",
		TimeStamp:  datetime.GetNowDateTime(),
		Version:    "1.0",
		BizContent: content,
	}

	signURL, err := query.Values(args)
	if err != nil {
		return nil, err
	}

	sign, err := NewToken().Sign(signURL, PrivatePath)
	if err != nil {
		return nil, err
	}

	signURL.Add("sign", sign)
	result, err := m.do(fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s",
		signURL.Encode()),
	)
	if err != nil {
		return nil, err
	}

	response := result.AlipayOpenAppMiniTemplateMessageSendResponse
	if response.Code != "10000" {
		return nil, errors.New(response.Msg)
	}
	return response, nil
}

// do message do
func (m *Message) do(url string) (*Result, error) {
	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: url,
		Method: "GET",
	})
	if err != nil {
		return nil, err
	}

	result := Result{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
