package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/tiantour/fetch"
	"github.com/tiantour/rsae"
)

var (
	// AppKey appkey
	AppKey string

	// AppSecret appsecret
	AppSecret string

	// Sign sign
	Sign string
)

type (
	// Request request
	Request struct {
		Method       string `json:"method" url:"method,omitempty"`                 // 是 接口名称
		AppKey       string `json:"app_key" url:"app_key,omitempty"`               // 是 app key
		TargetAppKey string `json:"target_app_key" url:"target_app_key,omitempty"` // 否 第三方 app key
		SignMethod   string `json:"sign_method" url:"sign_method,omitempty"`       // 是 签名 md5 hmac
		Sign         string `json:"sign" url:"sign,omitempty"`                     // 是 签名
		Session      string `json:"session" url:"session,omitempty"`               // 否 授权
		Timestamp    string `json:"timestamp" url:"timestamp,omitempty"`           // 是 时间 字符串
		Format       string `json:"format" url:"format,omitempty"`                 // 否 响应格式
		V            string `json:"v" url:"v,omitempty"`                           // 是 版本 2.0
		PartnerID    string `json:"partner_id" url:"partner_id,omitempty"`         // 否 合伙伙伴
		Simplify     string `json:"simplify" url:"simplify,omitempty"`             // 否 精简json
	}
	// Response response
	Response struct {
		AlibabaAliqinFcSmsNumSendResponse interface{} `json:"alibaba_aliqin_fc_sms_num_send_response"` // 正确
		ErrResponse                       interface{} `json:"err_response"`                            // 错误
	}
	// SMS sms
	SMS struct{}
)

// NewSMS new sms
func NewSMS() *SMS {
	return &SMS{}
}

// Body body
func (s *SMS) Body(args interface{}) ([]byte, error) {
	body, err := s.Sign(args)
	if err != nil {
		return []byte{}, err
	}
	header := http.Header{}
	header.Add("Content-Type", "application/x-www-form-urlencoded")
	result, err := fetch.Cmd(fetch.Request{
		Method: "POST",
		URL:    "https://eco.taobao.com/router/rest",
		Body:   []byte(body),
		Header: header,
	})
	if err != nil {
		return []byte{}, err
	}
	response := Response{}
	err = json.Unmarshal(result, &response)
	if err != nil || response.ErrResponse != nil {
		return []byte{}, errors.New(string(result))
	}
	return result, nil
}

// Sign sign
func (s *SMS) Sign(args interface{}) (string, error) {
	params, err := query.Values(args)
	if err != nil {
		return "", err
	}
	query, err := url.QueryUnescape(params.Encode())
	if err != nil {
		return "", err
	}
	sign := fmt.Sprintf("%s%s%s",
		AppSecret,
		query,
		AppSecret,
	)
	sign = strings.Replace(sign, "&", "", -1)
	sign = strings.Replace(sign, "=", "", -1)
	sign = strings.ToUpper(rsae.NewMD5().Encode(sign))
	return fmt.Sprintf("%s&sign=%s", query, sign), nil
}
