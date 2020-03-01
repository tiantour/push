package alipay

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/tiantour/fetch"
	"github.com/tiantour/imago"
	"github.com/tiantour/rsae"
	"github.com/tiantour/tempo"
)

// Token token
type Token struct{}

// NewToken new token
func NewToken() *Token {
	return &Token{}
}

// Access access token
func (t *Token) Access(code string) (*Response, error) {
	args := &Request{
		AppID:     AppID,
		Method:    "alipay.system.oauth.token",
		Format:    "JSON",
		Charset:   "utf-8",
		SignType:  "RSA2",
		TimeStamp: tempo.NewNow().String(),
		Version:   "1.0",
		GrantType: "authorization_code",
		Code:      code,
	}
	tmp, err := query.Values(args)
	if err != nil {
		return nil, err
	}
	sign, err := t.Sign(&tmp, PrivatePath)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s", sign)
	body, err := fetch.Cmd(&fetch.Request{
		Method: "GET",
		URL:    url,
	})
	if err != nil {
		return nil, err
	}
	result := Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	response := result.AlipaySystemOauthTokenResponse
	if response.Code != "10000" {
		return nil, errors.New(response.Msg)
	}
	return response, err
}

// Sign trade sign
func (t *Token) Sign(args *url.Values, privatePath string) (string, error) {
	query, err := url.QueryUnescape(args.Encode())
	if err != nil {
		return "", err
	}
	privateKey, err := imago.NewFile().Read(privatePath)
	if err != nil {
		return "", err
	}
	return rsae.NewRSA().Sign(query, privateKey)
}
