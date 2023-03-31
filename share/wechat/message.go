package wechat

import (
	"fmt"
	"net/url"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/duke-git/lancet/v2/random"
	"github.com/google/go-querystring/query"
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
		Noncestr:    random.RandString(16),
		Timestamp:   fmt.Sprintf("%d", datetime.NewUnixNow().ToUnix()),
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

	return cryptor.Sha1(query), nil
}
