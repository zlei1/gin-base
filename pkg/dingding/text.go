package dingding

import (
	"encoding/json"
)

type TextMessage struct {
	MsgType string      `json:"msgtype"`
	Text    TextContent `json:"text"`
	At      At          `json:"at,omitempty"`
}

func (msg *TextMessage) ToJsonByte() ([]byte, error) {
	msg.MsgType = MsgTypeText

	jsonByte, err := json.Marshal(msg)
	return jsonByte, err
}

type TextContent struct {
	Content string `json:"content"`
}

func NewTextMessage() *TextMessage {
	msg := &TextMessage{}
	return msg
}

func (msg *TextMessage) SetContent(content string) *TextMessage {
	msg.Text = TextContent{
		Content: content,
	}

	return msg
}

func (msg *TextMessage) SetAt(atMobiles []string, isAtAll bool) *TextMessage {
	msg.At = At{
		AtMobiles: atMobiles,
		IsAtAll:   isAtAll,
	}
	return msg
}
