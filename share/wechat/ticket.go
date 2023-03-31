package wechat

import (
	"errors"
	"fmt"
	"time"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/tiantour/push/x/cache"
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
	token, ok := cache.NewString().Get(AppID)
	if ok && token != "" {
		return token.(string), nil
	}

	result, err := t.Get()
	if err != nil {
		return "", err
	}

	_ = cache.NewString().Set(AppID, result.Ticket, 1, 7200*time.Second)
	return result.Ticket, nil
}

// Get get ticket
func (t *Ticket) Get() (*Ticket, error) {
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi", token),
		Method: "GET",
	})
	if err != nil {
		return nil, err
	}

	result := Ticket{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, err
}
