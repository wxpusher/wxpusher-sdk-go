package test

import (
	"github.com/TMaize/wxpusher-client"
	"github.com/TMaize/wxpusher-client/model"
	"testing"
)

func TestSendMessage(t *testing.T) {
	appToken := "AT_*****"
	uId := "UID_*****"
	msg := model.NewMessage(appToken).SetContent("测试消息").AddUId(uId)
	result, err := wxpusher.SendMessage(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}

func TestQueryMessageStatus(t *testing.T) {
	messageId := 2384250
	result, err := wxpusher.QueryMessageStatus(messageId)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}
