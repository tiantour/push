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
		AlibabaAliqinFcSmsNumSendResponse  Result `json:"alibaba_aliqin_fc_sms_num_send_response"`  // 正确
		AlibabaAliqinFcSmsNumQueryResponse Value  `json:"alibaba_aliqin_fc_sms_num_query_response"` //
		ErrResponse                        Fail   `json:"error_response"`                           // 错误
	}
	// Result result
	Result struct {
		Result Success `json:"result"` // 结果
	}

	// Success success
	Success struct {
		ErrCode string `json:"err_code"` // 返回错误
		Model   string `json:"model"`    // 返回结果
		Success bool   `json:"success"`  // 返回状态
		Msg     string `json:"message"`  // 返回描述
	}

	// Fail fail
	Fail struct {
		SubMsg  string `json:"sub_msg"`  // 错误信息
		SubCode string `json:"sub_code"` // 错误解释
		Code    int    `json:"code"`     // 错误代码
		Msg     string `json:"msg"`      // 错误描述
	}

	// Value value
	Value struct {
		CurrentPage int     `json:"current_page"` // 当前页码
		PageSize    int     `json:"page_size"`    // 每页数量
		TotalCount  int     `json:"total_count"`  // 总量
		TotalPage   int     `json:"total_page"`   // 总页数
		Values      Partner `json:"values"`
	}
	// Partner partner
	Partner struct {
		FcPartnerSmsDetailDto []Detail `json:"fc_partner_sms_detail_dto"`
	}
	// Detail detail
	Detail struct {
		Extend          string `json:"extend"`            // 公共回传参数
		RecNum          string `json:"rec_num"`           // 短信接收号码
		ResultCode      string `json:"result_code"`       // 短信错误码
		SMSCode         string `json:"sms_code"`          // 模板编码
		SMSContent      string `json:"sms_content"`       // 短信发送内容
		SMSReceiverTime string `json:"sms_receiver_time"` // 短信接收时间
		SMSSendTime     string `json:"sms_send_time"`     // 短信发送时间
		SMSStatus       int    `json:"sms_status"`        // 发送状态 1：等待回执，2：发送失败，3：发送成功
	}
)

// SMS sms
type SMS struct{}

// NewSMS new sms
func NewSMS() *SMS {
	return &SMS{}
}

// Body body
func (s *SMS) do(args interface{}) (*Response, error) {
	str, err := s.sign(args)
	if err != nil {
		return nil, err
	}
	body, err := fetch.Cmd(fetch.Request{
		Method: "POST",
		URL:    "https://eco.taobao.com/router/rest",
		Body:   []byte(str),
		Header: http.Header{
			"Content-Type": []string{"application/x-www-form-urlencoded"},
		},
	})
	if err != nil {
		return nil, err
	}
	result := Response{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrResponse.Code != 0 {
		return nil, errors.New(result.ErrResponse.Msg)
	}
	return &result, nil
}

// sign sign
func (s *SMS) sign(args interface{}) (string, error) {
	params, err := query.Values(args)
	if err != nil {
		return "", err
	}
	query, err := url.QueryUnescape(params.Encode())
	if err != nil {
		return "", err
	}
	sign := fmt.Sprintf("%s%s%s", AppSecret, query, AppSecret)
	sign = strings.Replace(sign, "&", "", -1)
	sign = strings.Replace(sign, "=", "", -1)
	sign = strings.ToUpper(rsae.NewMD5().Encode(sign))
	return fmt.Sprintf("%s&sign=%s", query, sign), nil
}
