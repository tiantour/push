package wechat

type (
	// UNI uni message
	UNI struct {
		ToUser           string  `json:"touser,omitempty"`                    // 是	接收者（用户）的 openid
		WeappTemplateMsg *Webapp `json:"weapp_template_msg,omitempty"`        // 否 小程序模板消息相关的信息，可以参考小程序模板消息接口; 有此节点则优先发送小程序模板消息
		MpTemplateMsg    *MP     `json:"mp_template_msg,omitempty,omitempty"` // 是 公众号模板消息相关的信息，可以参考公众号模板消息接口；有此节点并且没有weapp_template_msg节点时，发送公众号模板消息
	}

	// Webapp miniprogram message
	Webapp struct {
		TemplateID      string `json:"template_id,omitempty"`      // 是 小程序模板ID
		Page            string `json:"page,omitempty"`             // 是 小程序页面路径
		FormID          string `json:"form_id,omitempty"`          // 是 小程序模板消息formid
		Data            *Data  `json:"data,omitempty"`             // 是 小程序模板数据
		EmphasisKeyword string `json:"emphasis_keyword,omitempty"` // 是 小程序模板放大关键词
	}

	// MP office account message
	MP struct {
		AppID       string       `json:"appid,omitempty"`       // 是 公众号appid，要求与小程序有绑定且同主体
		TemplateID  string       `json:"template_id,omitempty"` // 是 公众号模板id
		URL         string       `json:"url,omitempty"`         // 是 公众号模板消息所要跳转的url
		MiniProgram *MiniProgram `json:"miniprogram,omitempty"` // 是 公众号模板消息所要跳转的小程序，小程序的必须与公众号具有绑定关系
		Data        *Data        `json:"data,omitempty"`        // 是 公众号模板消息的数据
	}

	// MI miniprogram message
	MI struct {
		ToUser           string `json:"touser,omitempty"`            // 是 接收者（用户）的 openid
		TemplateID       string `json:"template_id,omitempty"`       // 是 所需下发的订阅模板id
		Page             string `json:"page,omitempty"`              // 是 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
		Data             *Data  `json:"data,omitempty"`              // 是 小程序模板数据
		MiniprogramState string `json:"miniprogram_state,omitempty"` // 否 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
		Lang             string `json:"lang,omitempty"`              // 是 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
	}

	// Data data
	Data struct {
		First     *Option `json:"first,omitempty"`     // 标题
		Keyword1  *Option `json:"keyword1,omitempty"`  // 关键词1
		Keyword2  *Option `json:"keyword2,omitempty"`  // 关键词2
		Keyword3  *Option `json:"keyword3,omitempty"`  // 关键词3
		Keyword4  *Option `json:"keyword4,omitempty"`  // 关键词4
		Keyword5  *Option `json:"keyword5,omitempty"`  // 关键词5
		Keyword6  *Option `json:"keyword6,omitempty"`  // 关键词6
		Keyword7  *Option `json:"keyword7,omitempty"`  // 关键词7
		Keyword8  *Option `json:"keyword8,omitempty"`  // 关键词8
		Keyword9  *Option `json:"keyword9,omitempty"`  // 关键词9
		Keyword10 *Option `json:"keyword10,omitempty"` // 关键词10
		Remark    *Option `json:"remark,omitempty"`    // 备注
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
