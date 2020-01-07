package model

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// Qrcode 生成二维码参数
type Qrcode struct {
	AppToken  string `json:"appToken"`
	Extra     string `json:"extra"`
	ValidTime int    `json:"validTime"`
}

// CreateQrcodeResult 创建二维码返回参数
type CreateQrcodeResult struct {
	Expires  int    `json:"expires"`
	Code     string `json:"code"`
	ShortUrl string `json:"shortUrl"`
	Url      string `json:"url"`
	Extra    string `json:"extra"`
}

// Check参数校验
func (q *Qrcode) Check() error {
	if len(q.AppToken) == 0 {
		return NewBusinessError(errors.New("appToken不能为空"))
	}
	q.Extra = strings.TrimSpace(q.Extra)
	if q.Extra == "" {
		return NewBusinessError(errors.New("extra不能为空"))
	}
	if utf8.RuneCountInString(q.Extra) > 64 {
		return NewBusinessError(errors.New("extra长度不能超过64"))
	}
	if q.ValidTime > 2592000 {
		return NewBusinessError(errors.New("二维码有效期最大30天"))
	}
	// 默认值默认1800 s
	if q.ValidTime <= 0 {
		q.ValidTime = 1800
	}
	return nil
}
