package wechat

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/duke-git/lancet/v2/netutil"
)

// Message message
type Message struct {
	ErrCode int    `json:"errcode"` // 错误代码
	ErrMsg  string `json:"errmsg"`  // 错误消息
}

// NewMessage new message
func NewMessage() *Message {
	return &Message{}
}

// MI mi message
func (m *Message) MI(args *MI) (*Message, error) {
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}

	return m.do(args, fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s",
		token),
	)
}

// MP mp message
func (m *Message) MP(args *MP) (*Message, error) {
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", token)
	return m.do(args, url)
}

// do do message
func (m *Message) do(data interface{}, url string) (*Message, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: url,
		Method: "POST",
		Body:   body,
	})
	if err != nil {
		return nil, err
	}

	result := Message{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}
