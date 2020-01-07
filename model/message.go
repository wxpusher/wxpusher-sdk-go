package model

import (
	"errors"
)

// Message 消息
type Message struct {
	AppToken    string   `json:"appToken"`
	Content     string   `json:"content"`
	ContentType int      `json:"contentType"`
	TopicIds    []int    `json:"topicIds"`
	UIds        []string `json:"uids"`
	Url         string   `json:"url"`
}

// SendMsgResult 发送消息结果
type SendMsgResult struct {
	Uid       string `json:"uid"`
	TopicId   string `json:"topicId"`
	MessageId int    `json:"messageId"`
	Code      int    `json:"code"`
	Status    string `json:"status"`
}

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

func (m *Message) SetUrl(url string) *Message {
	m.Url = url
	return m
}

func (m *Message) AddUId(id string, more ...string) *Message {
	m.UIds = append(m.UIds, id)
	m.UIds = append(m.UIds, more...)
	return m
}

func (m *Message) AddTopicId(id int, more ...int) *Message {
	m.TopicIds = append(m.TopicIds, id)
	m.TopicIds = append(m.TopicIds, more...)
	return m
}

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
	if len(m.UIds) == 0 {
		return NewBusinessError(errors.New("必须添加uid"))
	}
	return nil
}
