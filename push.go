package push

import (
	"github.com/tiantour/push/notice"
	"github.com/tiantour/push/sms"
)

// IOS
var (
	SMS    = &smsx{}
	Notice = &noticex{}
)

type (
	// 发送短信
	smsx struct {
		Text  sms.Text  // 发送文字短信
		Voice sms.Voice // 发送语言短息
	}
	// 发送通知
	noticex struct {
		IOS     notice.IOS     // 发送苹果通知
		Android notice.Android // 发送安卓通知
		Wechat  notice.Wechat  // 发送微信通知
	}
)
