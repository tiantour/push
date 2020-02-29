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

	// next todo
	wechat.NewMessage().MI()
	wechat.NewMessage().MP()
	wechat.NewMessage().UNI()