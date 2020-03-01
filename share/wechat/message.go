package wechat

import (
	"encoding/hex"
	"fmt"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/tiantour/imago"
	"github.com/tiantour/rsae"
)

// Message message
type Message struct {
	AppID       string `json:"appid,omitempty" url:"-"`
	JSapiTicket string `json:"jsapi_ticket,omitempty" url:"jsapi_ticket,omitempty"`
	Noncestr    string `json:"noncestr,omitempty" url:"noncestr,omitempty"`
	Timestamp   string `json:"timestamp,omitempty" url:"timestamp,omitempty"`
	URL         string `json:"url,omitempty" url:"url,omitempty"`
	Signature   string `json:"signature,omitempty" url:"-"`
}

// NewMessage new message
func NewMessage() *Message {
	return &Message{}
}

// Share message share
func (m *Message) Share(path string) (*Message, error) {
	ticket, err := NewTicket().JSAPI()
	if err != nil {
		return nil, err
	}

	result := &Message{
		AppID:       AppID,
		Noncestr:    imago.NewRandom().Text(16),
		Timestamp:   fmt.Sprintf("%d", time.Now().Unix()),
		URL:         path,
		JSapiTicket: ticket,
	}
	signURL, err := query.Values(result)
	if err != nil {
		return nil, err
	}
	sign, err := m.sign(signURL)
	if err != nil {
		return nil, err
	}
	result.Signature = sign

	return result, nil
}

// sign
func (m *Message) sign(args url.Values) (string, error) {
	query, err := url.QueryUnescape(args.Encode())
	if err != nil {
		return "", err
	}
	body := rsae.NewSHA().SHA1(query)
	return hex.EncodeToString(body), nil
}
