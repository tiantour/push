package wechat

var (
	// AppID appid
	AppID string

	// AppSecret app secret
	AppSecret string
)

type (

	// MI miniprogram message
	MI struct {
		ToUser           string             `json:"touser,omitempty"`            // 是 接收者（用户）的 openid
		TemplateID       string             `json:"template_id,omitempty"`       // 是 所需下发的订阅模板id
		Page             string             `json:"page,omitempty"`              // 是 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
		Data             map[string]*Option `json:"data,omitempty"`              // 是 小程序模板数据
		MiniprogramState string             `json:"miniprogram_state,omitempty"` // 否 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
		Lang             string             `json:"lang,omitempty"`              // 是 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
	}

	// MP office account message
	MP struct {
		ToUser      string             `json:"touser,omitempty"`      // 是 公众号appid，要求与小程序有绑定且同主体
		TemplateID  string             `json:"template_id,omitempty"` // 是 公众号模板id
		URL         string             `json:"url,omitempty"`         // 是 公众号模板消息所要跳转的url
		MiniProgram *MiniProgram       `json:"miniprogram,omitempty"` // 是 公众号模板消息所要跳转的小程序，小程序的必须与公众号具有绑定关系
		Data        map[string]*Option `json:"data,omitempty"`        // 是 公众号模板消息的数据
	}

	// Option option
	Option struct {
		Value string `json:"value,omitempty"` // 内容
		Color string `json:"color,omitempty"` // 颜色
	}

	// MiniProgram  miniprogram
	MiniProgram struct {
		AppID    string `json:"appid,omitempty"`    // 小程序编号
		PagePath string `json:"pagepath,omitempty"` // 小程序路径
	}
)
