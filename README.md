# sms message

aliyun

```
package main

import (
	"fmt"

	"github.com/tiantour/push/sms/aliyun"
)

func main() {
	aliyun.AccessKeyID = "yourAccessKeyID"
	aliyun.AccessKeySecret = "your AccessKeySecret"
	aliyun.Sign = "your sms sign"

	query()
    send()
}

func query() {
	x, err := aliyun.NewMessage().Query(&aliyun.Message{
		Phone: []string{"your phone number"},
		Date:  "2020-02-28",
		Page:  0,
		Size:  10,
	})
	fmt.Println(x, err)
}

func send() {
	x, err := aliyun.NewMessage().Send(&aliyun.Message{
		Phone:    []string{"your phone number"},
		Template: "your template id",
		Body: map[string]string{
			"code": "123456",
		},
	})
	fmt.Println(x, err)
}
```

qiniu

```
package main

import (
	"fmt"

	"github.com/tiantour/push/sms/qiniu"
)

func main() {
	qiniu.AccessKey = "your AccessKey"
	qiniu.SecretKey = "your SecretKey"
	
    query()
	send()
}

func query() {
	x, err := qiniu.NewMessage().Query(&qiniu.Message{
		Phone: []string{"your phone number"},
		Date:  "2020-02-28",
		Page:  0,
		Size:  10,
	})
	fmt.Println(x, err)
}
func send() {
	x, err := qiniu.NewMessage().Send(&qiniu.Message{
		Phone:    []string{"your phoe number"},
		Template: "your template id",
		Body: map[string]string{
			"code": "123456",
		},
	})
	fmt.Println(x, err)
}

```

# template message

alipay

	// next todo
	alipay.NewMessage().MI()

wechat

	```
	package main

	import (
		"fmt"

		"github.com/tiantour/push/template/wechat"
	)

	func main() {
		mi()
		mp()
	}

	func mi() {
		wechat.AppID = "your AppID"
		wechat.AppSecret = "your AppSecret"

		data := map[string]*wechat.Option{
			"first": &wechat.Option{
				Value: "恭喜你发布成功",
			},
			"keyword1": &wechat.Option{
				Value: "发布文章",
			},
			"keyword2": &wechat.Option{
				Value: "2020-03-01",
			},
			"remark": &wechat.Option{
				Value: "欢迎再次购买",
			},
		}
		result, err := wechat.NewMessage().MP(&wechat.MP{
			ToUser:     "your openid",
			TemplateID: "your template id",
			URL:        "your page url",
			Data:       data,
		})
		fmt.Println(result, err)
	}
	
	func mp() {
		data := map[string]*wechat.Option{
			"thing1": &wechat.Option{
				Value: "test",
			},
			"amount2": &wechat.Option{
				Value: "￥0.01",
			},
			"date3": &wechat.Option{
				Value: "2020-03-01",
			},
			"thing4": &wechat.Option{
				Value: "remark",
			},
		}
		result, err := wechat.NewMessage().MI(&wechat.MI{
			ToUser:     "your openid",
			TemplateID: "your template id",
			Page:       "your page path",
			Data:       data,
		})
		fmt.Println(result, err)
	}
	```