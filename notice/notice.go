package notice

type (
	// IOS ios
	IOS struct{}
	// Android android
	Android struct{}
	// Wechat wechat
	Wechat struct {
		AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
		ExpiresIn    int    `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
		RefreshToken string `json:"refresh_token"` // 用户刷新access_token
		OpenID       string `json:"openid"`        // 用户唯一标识
		Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
		ErrCode      int    `json:"errcode"`       // 错误代码
		ErrMsg       string `json:"errmsg"`        // 错误消息
	}
)
