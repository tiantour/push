package wechat

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/cache"
	"github.com/tiantour/fetch"
)

// Ticket ticket
type Ticket struct {
	Ticket    string `json:"ticket,omitempty"`
	ExpiresIn int    `json:"expires_in,omitempty"`
	ErrCode   int    `json:"errcode,omitempty"`
	ErrMsg    string `json:"errmsg,omitempty"`
}

// NewTicket new ticket
func NewTicket() *Ticket {
	return &Ticket{}
}

// JSAPI jsapi ticket
func (t *Ticket) JSAPI() (string, error) {
	result, err := t.Cache()
	if err != nil {
		ticket, err := t.Network()
		if err != nil {
			return "", err
		}

		result = ticket.Ticket
		key := fmt.Sprintf("string:data:bind:jsapi:ticket:%s", AppID)
		_ = cache.NewString().SET(nil, key, result, "EX", 7200)
	}
	return result, nil
}

// Cache get ticket from cache
func (t *Ticket) Cache() (string, error) {
	var result string
	key := fmt.Sprintf("string:data:bind:jsapi:ticket:%s", AppID)
	err := cache.NewString().GET(&result, key)
	return result, err
}

// Network get ticket form network
func (t *Ticket) Network() (*Ticket, error) {
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}
	body, err := fetch.Cmd(&fetch.Request{
		Method: "GET",
		URL: fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi",
			token,
		),
	})
	if err != nil {
		return nil, err
	}
	result := Ticket{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, err
}
