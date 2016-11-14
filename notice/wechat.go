package notice

import (
	"encoding/json"
	"fmt"

	"github.com/tiantour/cache"
	"github.com/tiantour/fetch"
)

// Token message
func (w *Wechat) token(appID, appSecret string) (string, error) {
	key := "wechat_access_token"
	token, err := cache.String.GET(key).Str()
	if err != nil {
		result := Wechat{}
		url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
			appID,
			appSecret,
		)
		body, err := w.request(url)
		if err == nil && json.Unmarshal(body, &result) == nil {
			_ = cache.String.SET(key, result.AccessToken)
			_ = cache.Key.EXPIRE(key, 7200)
			return result.AccessToken, nil
		}
		return "", err
	}
	return token, nil

}

// Push message
func (w *Wechat) Push(appID, appSecret string, data interface{}) ([]byte, error) {
	accessToken, err := w.token(appID, appSecret)
	if err != nil {
		return nil, err
	}
	requestURL := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken)
	requestData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return fetch.Cmd("post", requestURL, requestData)
}

// request
func (w *Wechat) request(requestURL string) ([]byte, error) {
	return fetch.Cmd("get", requestURL)
}
