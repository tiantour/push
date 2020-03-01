package wechat

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/fetch"
)

var (
	// AppID appid
	AppID string

	// AppSecret app secret
	AppSecret string
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

// UNI uni message
func (m *Message) UNI(args *UNI) (*Message, error) {
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}
	return m.do(args, fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=%s",
		token),
	)
}

// do
func (m *Message) do(data interface{}, url string) (*Message, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body, err = fetch.Cmd(&fetch.Request{
		Method: "POST",
		URL:    url,
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	result := Message{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}
