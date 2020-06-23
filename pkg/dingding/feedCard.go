package dingding

import (
	"encoding/json"
)

type FeedCardMessage struct {
	MsgType  string          `json:"msgtype"`
	FeedCard FeedCardContent `json:"feedCard"`
}

func (msg *FeedCardMessage) ToJsonByte() ([]byte, error) {
	msg.MsgType = MsgTypeFeedCard

	jsonByte, err := json.Marshal(msg)
	return jsonByte, err
}

type FeedCardContent struct {
	Links []FeedCardContentLink `json:"links"`
}

type FeedCardContentLink struct {
	Title      string `json:"title"`
	PicURL     string `json:"picURL"`
	MessageURL string `json:"messageURL"`
}

func NewFeedCardMessage() *FeedCardMessage {
	msg := &FeedCardMessage{}
	return msg
}

func (msg *FeedCardMessage) AddContent(title string, messageURL string, picURL string) *FeedCardMessage {
	link := FeedCardContentLink{
		Title:      title,
		MessageURL: messageURL,
		PicURL:     picURL,
	}

	msg.FeedCard.Links = append(msg.FeedCard.Links, link)
	return msg
}
