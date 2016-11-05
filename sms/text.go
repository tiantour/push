package sms

import (
	"encoding/json"

	"github.com/ltt1987/alidayu"
)

// Send sms
func (t *Text) Send(number, sign, template string, param map[string]string) (bool, string) {
	paramByte, _ := json.Marshal(param)
	paramString := string(paramByte)
	return alidayu.SendSMS(
		number,
		sign,
		template,
		paramString,
	)
}
