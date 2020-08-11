# wxpusher-client

[zjiecode/wxpusher-client](https://github.com/zjiecode/wxpusher-client)的 Go 语言版本

## 安装

```sh
go get -u github.com/wxpusher/wxpusher-sdk-go
```

```sh
go get github.com/wxpusher/wxpusher-sdk-go@v1.0.3
```

引入

```go
import (
	"github.com/wxpusher/wxpusher-sdk-go"
	"github.com/wxpusher/wxpusher-sdk-go/model"
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

## 查询 App 的关注用户

```go
result, err := wxpusher.QueryWxUser(appToken, 1, 20)
fmt.Println(result, err)
```
