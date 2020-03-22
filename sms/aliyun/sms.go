package aliyun

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/tiantour/rsae"
)

// SMS sms
type SMS struct{}

// NewSMS new sms
func NewSMS() *SMS {
	return &SMS{}
}

// Sign sms sign
func (s *SMS) Sign(args url.Values) string {
	query := args.Encode()
	query = strings.Replace(query, "+", "%20", -1)
	query = strings.Replace(query, "*", "%2A", -1)
	query = strings.Replace(query, "%7E", "~", -1)
	query = fmt.Sprintf("POST&%s&%s", url.QueryEscape("/"), url.QueryEscape(query))

	body := rsae.NewSHA().HmacSha1(query, AccessKeySecret+"&")
	return rsae.NewBase64().Encode(body)
}
