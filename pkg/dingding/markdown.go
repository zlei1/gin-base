package dingding

import (
	"encoding/json"
)

type MarkdownMessage struct {
	MsgType  string          `json:"msgtype"`
	Markdown MarkdownContent `json:"markdown"`
	At       At              `json:"at,omitempty"`
}

func (msg *MarkdownMessage) ToJsonByte() ([]byte, error) {
	msg.MsgType = MsgTypeMarkdown

	jsonByte, err := json.Marshal(msg)
	return jsonByte, err
}

type MarkdownContent struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func NewMarkdownMessage() *MarkdownMessage {
	msg := &MarkdownMessage{}
	return msg
}

func (msg *MarkdownMessage) SetContent(title string, text string) *MarkdownMessage {
	msg.Markdown = MarkdownContent{
		Title: title,
		Text:  text,
	}

	return msg
}

func (msg *MarkdownMessage) SetAt(atMobiles []string, isAtAll bool) *MarkdownMessage {
	msg.At = At{
		AtMobiles: atMobiles,
		IsAtAll:   isAtAll,
	}
	return msg
}
