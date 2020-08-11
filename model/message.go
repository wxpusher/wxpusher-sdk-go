package model

import (
	"errors"
)

// Message 消息
//  ContentType： 1-文字 2-html
type Message struct {
	AppToken    string   `json:"appToken"`
	Content     string   `json:"content"`
	Summary     string   `json:"summary"`
	ContentType int      `json:"contentType"`
	TopicIds    []int    `json:"topicIds"`
	UIds        []string `json:"uids"`
	Url         string   `json:"url"`
}

// SendMsgResult 发送消息接口的响应结果
type SendMsgResult struct {
	Uid       string `json:"uid"`
	TopicId   int    `json:"topicId"`
	MessageId int    `json:"messageId"`
	Code      int    `json:"code"`
	Status    string `json:"status"`
}

// NewMessage 使用AppToken创建消息
func NewMessage(appToken string) *Message {
	return &Message{
		AppToken:    appToken,
		ContentType: 1,
	}
}

func (m *Message) SetContentType(ct int) *Message {
	m.ContentType = ct
	return m
}

func (m *Message) SetContent(content string) *Message {
	m.Content = content
	return m
}

func (m *Message) SetSummary(summary string) *Message {
	m.Summary = summary
	return m
}

func (m *Message) SetUrl(url string) *Message {
	m.Url = url
	return m
}

func (m *Message) AddUId(id string, more ...string) *Message {
	m.UIds = append(m.UIds, id)
	m.UIds = append(m.UIds, more...)
	return m
}

// AddTopicId 添加主题
func (m *Message) AddTopicId(id int, more ...int) *Message {
	m.TopicIds = append(m.TopicIds, id)
	m.TopicIds = append(m.TopicIds, more...)
	return m
}

// Check 本地校验消息的正确性
func (m Message) Check() error {
	if len(m.AppToken) == 0 {
		return NewBusinessError(errors.New("appToken不能为空"))
	}
	if m.ContentType == 0 {
		return NewBusinessError(errors.New("必须指定contentType"))
	}
	if m.ContentType < 0 || m.ContentType > 3 {
		return NewBusinessError(errors.New("错误的contentType"))
	}
	if len(m.Content) == 0 {
		return NewBusinessError(errors.New("content内容不能为空"))
	}
	if len(m.UIds)+len(m.TopicIds) == 0 {
		return NewBusinessError(errors.New("未指定接收者，请设置UIds/TopicIds"))
	}
	return nil
}
