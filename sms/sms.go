package sms

import (
	"github.com/ltt1987/alidayu"
	"github.com/tiantour/conf"
)

type (
	// Text text
	Text struct{}
	// Voice text
	Voice struct{}
)

func init() {
	alidayu.AppKey = conf.Options.Alidayu.AppKey
	alidayu.AppSecret = conf.Options.Alidayu.AppSecret
}
