# alidayu
alidayu golang package

### SMS

send

```
package main

import (
	"fmt"

	"github.com/tiantour/alidayu/sms"
)

func init() {
	sms.AppKey = "xxx"
	sms.AppSecret = "xxx"
	sms.Sign = "xxx"
}
func main() {
	x, err := sms.NewSend().Do(
		"18910315767",
		"SMS_10271157",
		map[string]string{
			"number": "123456",
		})
	fmt.Println(x, err)
}
```

query

```
package main

import (
	"fmt"

	"github.com/tiantour/alidayu/sms"
)

func init() {
	sms.AppKey = "xxx"
	sms.AppSecret = "xxx"
	sms.Sign = "xxx"
}
func main() {
	x, err := sms.NewQuery().Do(
		"18910315767",
		"20170524",
		1,
		10,
	)
	fmt.Println(x, err)
}
```

