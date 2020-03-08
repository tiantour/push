package wechat

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/cache"
	"github.com/tiantour/fetch"
)

// Token token
type Token struct {
	AccessToken  string `json:"access_token,omitempty"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int    `json:"expires_in,omitempty"`    // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token,omitempty"` // 用户刷新access_token
	OpenID       string `json:"openid,omitempty"`        // 用户唯一标识
	Scope        string `json:"scope,omitempty"`         // 用户授权的作用域，使用逗号（,）分隔
	ErrCode      int    `json:"errcode,omitempty"`       // 错误代码
	ErrMsg       string `json:"errmsg,omitempty"`        // 错误消息
}

// NewToken new token
func NewToken() *Token {
	return &Token{}
}

// Access access token
func (t *Token) Access() (string, error) {
	result, err := t.Cache()
	if err != nil || result == "" {
		token, err := t.Network()
		if err != nil {
			return "", err
		}

		result = token.AccessToken
		key := fmt.Sprintf("string:data:bind:access:token:%s", AppID)
		_ = cache.NewString().SET(nil, key, result, "EX", 7200)
	}
	return result, nil
}

// Cache get token from cache
func (t *Token) Cache() (string, error) {
	var result string
	key := fmt.Sprintf("string:data:bind:access:token:%s", AppID)
	err := cache.NewString().GET(&result, key)
	return result, err
}

// Network get token from network
func (t *Token) Network() (*Token, error) {
	body, err := fetch.Cmd(&fetch.Request{
		Method: "GET",
		URL: fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
			AppID,
			AppSecret,
		),
	})
	if err != nil {
		return nil, err
	}

	result := Token{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, err
}
