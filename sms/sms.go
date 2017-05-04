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
	alidayu.AppKey = conf.Data.Alidayu.AppKey
	alidayu.AppSecret = conf.Data.Alidayu.AppSecret
}
