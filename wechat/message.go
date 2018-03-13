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

// OA oa
func (m *Message) OA(body []byte) (*Message, error) {
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", token)
	return m.do(body, url)
}

// MP mp
func (m *Message) MP(body []byte) (*Message, error) {
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=%s", token)
	return m.do(body, url)
}

// do
func (m *Message) do(body []byte, url string) (*Message, error) {
	result := Message{}
	body, err := fetch.Cmd(fetch.Request{
		Method: "POST",
		URL:    url,
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}
