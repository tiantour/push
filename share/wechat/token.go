package wechat

import (
	"errors"
	"fmt"
	"time"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/tiantour/push/x/cache"
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
	token, ok := cache.NewString().Get(AppID)
	if ok && token != "" {
		return token.(string), nil
	}

	result, err := t.Get()
	if err != nil {
		return "", err
	}

	_ = cache.NewString().Set(AppID, result.AccessToken, 1, 7200*time.Second)
	return result.AccessToken, nil
}

// Get get token
func (t *Token) Get() (*Token, error) {
	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", AppID, AppSecret),
		Method: "GET",
	})
	if err != nil {
		return nil, err
	}

	result := Token{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, err
}
