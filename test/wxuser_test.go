package test

import (
	"github.com/wxpusher/wxpusher-sdk-go"
	"testing"
)

func TestQueryWxUser(t *testing.T) {
	appToken := "AT_*****"
	result, err := wxpusher.QueryWxUser(appToken, 1, 20)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}
