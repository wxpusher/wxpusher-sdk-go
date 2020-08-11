package test

import (
	"os"
	"testing"

	"github.com/wxpusher/wxpusher-sdk-go"
	"github.com/wxpusher/wxpusher-sdk-go/model"
)

func TestSendMessage(t *testing.T) {
	msg := model.NewMessage(os.Getenv("appToken")).SetContent("测试消息").SetSummary("这是摘要").AddUId(os.Getenv("UID"))
	result, err := wxpusher.SendMessage(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}

func TestSendTopic(t *testing.T) {
	msg := model.NewMessage(os.Getenv("appToken")).SetContent("测试消息").SetSummary("这是摘要").AddTopicId(640)
	result, err := wxpusher.SendMessage(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}

func TestQueryMessageStatus(t *testing.T) {
	messageID := 16727048
	result, err := wxpusher.QueryMessageStatus(messageID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}

func TestCreateQrcode(t *testing.T) {
	qrcode := model.Qrcode{
		AppToken: os.Getenv("appToken"),
		Extra:    "参数",
	}
	result, err := wxpusher.CreateQrcode(&qrcode)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}

func TestQueryWxUser(t *testing.T) {
	result, err := wxpusher.QueryWxUser(os.Getenv("appToken"), 1, 20)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}
