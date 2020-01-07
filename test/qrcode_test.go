package test

import (
	"github.com/TMaize/wxpusher-client"
	"github.com/TMaize/wxpusher-client/model"
	"testing"
)

func TestCreateQrcode(t *testing.T) {
	appToken := "AT_*****"
	qrcode := model.Qrcode{
		AppToken: appToken,
		Extra:    "参数",
	}
	result, err := wxpusher.CreateQrcode(&qrcode)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}
