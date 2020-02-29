package qiniu

var (
	// AccessKey access key
	AccessKey string

	// SecretKey secret key
	SecretKey string
)

type (
	// QueryRequest query request
	QueryRequest struct {
		JobID      string `json:"job_id,omitempty"`      // 否 发送任务返回的 id
		MessageID  string `json:"message_id,omitempty"`  // 否 单条短信发送接口返回的 id
		Mobile     string `json:"mobile,omitempty"`      // 否 接收短信的手机号码
		Status     string `json:"status,omitempty"`      // 否 短信的状态，sending: 发送中，success: 发送成功，failed: 发送失败，waiting: 等待发送
		TemplateID string `json:"template_id,omitempty"` // 否 模版 id
		Type       string `json:"type,omitempty"`        // 否 短信类型，marketing: 营销短信，notification: 通知短信，verification: 验证码类短信，voice: 语音短信
		Start      string `json:"start,omitempty"`       // 否 开始时间，timestamp，例如: 1563280448
		End        string `json:"end,omitempty"`         // 否 结束时间，timestamp，例如: 1563280471
		Page       int    `json:"page,omitempty"`        // 否 页码，默认为 1
		PageSize   int    `json:"page_size,omitempty"`   // 否 每页返回的数据条数，默认20，最大200
	}

	// QueryResponse query respose
	QueryResponse struct {
		Page     int             `json:"page,omitempty"`      // 否	页码，默认为 1
		PageSize int             `json:"page_size,omitempty"` // 否	每页返回的数据条数，默认20，最大200
		Total    int             `json:"total,omitempty"`     // 否	数量
		Items    []*QueryMessage `json:"items,omitempty"`     // 否	内容
	}

	// QueryMessage query message
	QueryMessage struct {
		JobID     string `json:"job_id,omitempty"`     // 否 发送任务返回的 id
		MessageID string `json:"message_id,omitempty"` // 否 单条短信发送接口返回的 id
		Mobile    string `json:"mobile,omitempty"`     // 否 接收短信的手机号码
		Content   string `json:"content,omitempty"`    // 否 短信内容
		Status    string `json:"status,omitempty"`     // 否 短信的状态，sending: 发送中，success: 发送成功，failed: 发送失败，waiting: 等待发送
		Type      string `json:"type,omitempty"`       // 否 短信类型，marketing: 营销短信，notification: 通知短信，verification: 验证码类短信，voice: 语音短信
		Error     string `json:"error,omitempty"`      // 否 短信发送失败详细状态信息
		Count     int    `json:"count,omitempty"`      // 否 发送的短信条数
		CreatedAt int    `json:"created_at,omitempty"` // 否 短信发送时间
		DelivrdAt int    `json:"delivrd_at,omitempty"` // 否 如果短信发送失败，那么不会返回这个字段
	}
	// SendRequest send response
	SendRequest struct {
		TemplateID string            `json:"template_id,omitempty"` // 是 模板 ID
		Mobiles    []string          `json:"mobiles,omitempty"`     // 是 手机号
		Parameters map[string]string `json:"parameters,omitempty"`  // 否 参数
	}
	// SendResponse send response
	SendResponse struct {
		JobID     string `json:"job_id,omitempty"`     // 否 发送任务返回的 id
		MessageID string `json:"message_id,omitempty"` // 否 单条短信发送接口返回的 id
	}
)
