package dingding

import (
	"encoding/json"
)

type LinkMessage struct {
	MsgType string      `json:"msgtype"`
	Link    LinkContent `json:"link"`
}

func (msg *LinkMessage) ToJsonByte() ([]byte, error) {
	msg.MsgType = MsgTypeLink

	jsonByte, err := json.Marshal(msg)
	return jsonByte, err
}

type LinkContent struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	PicURL     string `json:"picUrl"`
	MessageURL string `json:"messageUrl"`
}

func NewLinkMessage() *LinkMessage {
	msg := &LinkMessage{}
	return msg
}

func (msg *LinkMessage) SetContent(title string, text string, picURL string, messageURL string) *LinkMessage {
	msg.Link = LinkContent{
		Title:      title,
		Text:       text,
		PicURL:     picURL,
		MessageURL: messageURL,
	}

	return msg
}
