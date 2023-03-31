package qiniu

import (
	"net/url"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/netutil"
)

// SMS sms
type SMS struct{}

// NewSMS new sms
func NewSMS() *SMS {
	return &SMS{}
}

// Sign sms sign
func (s *SMS) Sign(args *netutil.HttpRequest) (string, error) {
	u, err := url.Parse(args.RawURL)
	if err != nil {
		return "", err
	}

	// 1. 添加 Path
	data := args.Method + " " + u.Path

	// 2. 添加 Query，前提: Query 存在且不为空
	if u.RawQuery != "" {
		data += "?" + u.RawQuery
	}

	// 3. 添加 Host
	data += "\nHost: " + u.Host

	// 4. 添加 Content-Type，前提: Content-Type 存在且不为空
	ctType := args.Headers.Get("Content-Type")
	if ctType != "" {
		data += "\nContent-Type: " + ctType
	}

	// 5. 添加回车
	data += ("\n\n")

	// 6. 添加 Body，前提: Content-Length 存在且 Body 不为空，
	// 同时 Content-Type 存在且不为空或 "application/octet-stream"
	ctLength := args.Headers.Get("Content-Length")
	if ctLength != "" && len(args.Body) != 0 {
		if ctType != "" && ctType != "application/octet-stream" {
			data += string(args.Body)
		}
	}

	str := cryptor.HmacSha1(data, SecretKey)
	return cryptor.Base64StdEncode(str), nil
}
