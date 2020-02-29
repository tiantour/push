package aliyun

var (
	// AccessKeyID AccessKeyID
	AccessKeyID string

	// AccessKeySecret AccessKeySecret
	AccessKeySecret string

	// Sign sign
	Sign string
)

type (
	// Request request
	Request struct {
		Format           string `json:"Format,omitempty" url:"Format,omitempty"`                     // 是 格式 json或xml
		RegionID         string `json:"RegionId,omitempty" url:"RegionId,omitempty"`                 // 是 地区 默认cn-hangzhou
		SignatureNonce   string `json:"SignatureNonce,omitempty" url:"SignatureNonce,omitempty"`     // 是 签名随机 32位随机数
		SignatureMethod  string `json:"SignatureMethod,omitempty" url:"SignatureMethod,omitempty"`   // 是 签名算法 默认HMAC-SHA1
		SignatureVersion string `json:"SignatureVersion,omitempty" url:"SignatureVersion,omitempty"` // 是 签名版本 默认1.0
		Signature        string `json:"Signature,omitempty" url:"Signature,omitempty"`               // 是 签名
		Timestamp        string `json:"Timestamp,omitempty" url:"Timestamp,omitempty"`               // 是 时间戳 默认标准时间，东八区-8小时
		Version          string `json:"Version,omitempty" url:"Version,omitempty"`                   // 是 版本号 默认2017-05-25
	}

	// Response response
	Response struct {
		Code      string `json:"Code,omitempty" url:"Code,omitempty"`           // OK	请求状态码。
		Message   string `json:"Message,omitempty" url:"Message,omitempty"`     // OK	状态码的描述。
		RequestID string `json:"RequestId,omitempty" url:"RequestId,omitempty"` // F655A8D5-B967-440B-8683-DAD6FF8DE990 请求ID。
	}

	// QueryRequest query request
	QueryRequest struct {
		CurrentPage int    `json:"CurrentPage,omitempty" url:"CurrentPage,omitempty"` // 是 1	分页查看发送记录，指定发送记录的的当前页码。
		PageSize    int    `json:"PageSize,omitempty" url:"PageSize,omitempty"`       // 是 10	分页查看发送记录，指定每页显示的短信记录数量。
		PhoneNumber string `json:"PhoneNumber,omitempty" url:"PhoneNumber,omitempty"` // 是 15900000000	接收短信的手机号码。
		SendDate    string `json:"SendDate,omitempty" url:"SendDate,omitempty"`       // 是 20181228 短信发送日期，支持查询最近30天的记录。格式为yyyyMMdd，例如20181225。
		AccessKeyID string `json:"AccessKeyId,omitempty" url:"AccessKeyId,omitempty"` // 否	LTAIP00vvvvvvvvv 主账号AccessKey的ID，即Key。
		Action      string `json:"Action,omitempty" url:"Action,omitempty"`           // 否	QuerySendDetails 系统规定参数。取值：QuerySendDetails。
		BizID       string `json:"BizId,omitempty" url:"BizId,omitempty"`             // 否	134523^4351232 发送回执ID，即发送流水号。调用发送接口SendSms或SendBatchSms发送短信时，返回值中的BizId字段。
		Request
	}

	// QueryResponse query response
	QueryResponse struct {
		SmsSendDetailDTOs QueryResponseList `json:"SmsSendDetailDTOs,omitempty" url:"SmsSendDetailDTOs,omitempty"` // 短信发送明细。
		TotalCount        int               `json:"TotalCount,omitempty" url:"TotalCount,omitempty"`               // 1 短信发送总条数。
		Response
	}

	// QueryResponseList query response list
	QueryResponseList struct {
		SmsSendDetailDTO []*QueryResponseItem `json:"SmsSendDetailDTO,omitempty"` // 短信发送明细。
	}

	// QueryResponseItem query  response item
	QueryResponseItem struct {
		Content      string `json:"Content,omitempty"`      //【阿里云】验证码为：123，您正在登录，若非本人操作，请勿泄露短信内容。
		ErrCode      string `json:"ErrCode,omitempty"`      // DELIVERED	运营商短信状态码。短信发送成功：DELIVERED 短信发送失败：失败错误码请参考错误码文档。
		OutID        string `json:"OutId,omitempty"`        // 123	外部流水扩展字段。
		PhoneNum     string `json:"PhoneNum,omitempty"`     // 15200000000 接收短信的手机号码。
		ReceiveDate  string `json:"ReceiveDate,omitempty"`  // 2019-01-08 16:44:13	短信接收日期和时间。
		SendDate     string `json:"SendDate,omitempty"`     // 2019-01-08 16:44:10 短信发送日期和时间。
		SendStatus   int    `json:"SendStatus,omitempty"`   // 3 短信发送状态，包括：1：等待回执。2：发送失败。3：发送成功。
		TemplateCode string `json:"TemplateCode,omitempty"` // SMS_122310183 短信模板ID。
	}

	// SendRequest send request
	SendRequest struct {
		AccessKeyID     string `json:"AccessKeyId,omitempty" url:"AccessKeyId,omitempty"`         // 否	LTAIP00vvvvvvvvv 主账号AccessKey的ID。
		Action          string `json:"Action,omitempty" url:"Action,omitempty"`                   // 否	SendSms	系统规定参数。取值：SendSms。
		OutID           string `json:"OutId,omitempty" url:"OutId,omitempty"`                     // 否	abcdefgh	外部流水扩展字段。
		PhoneNumbers    string `json:"PhoneNumbers,omitempty" url:"PhoneNumbers,omitempty"`       // 是 15900000000 接收短信的手机号码。
		SignName        string `json:"SignName,omitempty" url:"SignName,omitempty"`               // 是	阿里云	短信签名名称。请在控制台签名管理页面签名名称一列查看。
		SmsUpExtendCode string `json:"SmsUpExtendCode,omitempty" url:"SmsUpExtendCode,omitempty"` // 否	90999 上行短信扩展码，无特殊需要此字段的用户请忽略此字段。
		TemplateCode    string `json:"TemplateCode,omitempty" url:"TemplateCode,omitempty"`       // 是 SMS_153055065	短信模板ID。请在控制台模板管理页面模板CODE一列查看。
		TemplateParam   string `json:"TemplateParam,omitempty" url:"TemplateParam,omitempty"`     // 否 {"code":"1111"}	短信模板变量对应的实际值，JSON格式。
		Request
	}

	// SendResponse send response
	SendResponse struct {
		BizID string `json:"BizId,omitempty"` // 900619746936498440^0 发送回执ID，可根据该ID在接口QuerySendDetails中查询具体的发送状态。
		Response
	}
)
