# wxpusher-client

[zjiecode/wxpusher-client](https://github.com/zjiecode/wxpusher-client)的 Go 语言版本

## 安装

```sh
go get -u github.com/TMaize/wxpusher-client
```

引入

```go
import (
	"github.com/TMaize/wxpusher-client"
	"github.com/TMaize/wxpusher-client/model"
)
```

## 发送消息

```go
msg := model.NewMessage(appToken).SetContent("测试").AddUId(uId)
msgArr, err := wxpusher.SendMessage(msg)
fmt.Println(msgArr, err)
```

## 查询状态

```go
status, err := wxpusher.QueryMessageStatus(2384429)
fmt.Println(status, err)
```

## 创建参数二维码

```go
qrcode := model.Qrcode{AppToken: appToken, Extra: "XX渠道用户"}
qrcodeResp, err := wxpusher.CreateQrcode(&qrcode)
fmt.Println(qrcodeResp, err)
```

## 查询App的关注用户

```go
result, err := wxpusher.QueryWxUser(appToken, 1, 20)
fmt.Println(result, err)
```