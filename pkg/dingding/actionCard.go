package dingding

import (
	"encoding/json"
)

type ActionCardMessage struct {
	MsgType    string            `json:"msgtype"`
	ActionCard ActionCardContent `json:"actionCard"`
}

func (msg *ActionCardMessage) ToJsonByte() ([]byte, error) {
	msg.MsgType = MsgTypeActionCard

	jsonByte, err := json.Marshal(msg)
	return jsonByte, err
}

type ActionCardContent struct {
	Title          string                 `json:"title"`
	Text           string                 `json:"text"`
	SingleTitle    string                 `json:"singleTitle"`
	SingleURL      string                 `json:"singleURL"`
	Btns           []ActionCardContentBtn `json:"btns"`
	BtnOrientation string                 `json:"btnOrientation"`
}

type ActionCardContentBtn struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

func NewActionCardMessage() *ActionCardMessage {
	msg := &ActionCardMessage{}
	return msg
}

func (msg *ActionCardMessage) SetContentOverAll(
	title string,
	text string,
	btnOrientation string,
	singleTitle string,
	singleURL string) *ActionCardMessage {
	msg.ActionCard = ActionCardContent{
		Title:          title,
		Text:           text,
		BtnOrientation: btnOrientation,
		SingleTitle:    singleTitle,
		SingleURL:      singleURL,
	}
	return msg
}

func (msg *ActionCardMessage) SetContentIndependent(
	title string,
	text string,
	btns []ActionCardContentBtn,
	btnOrientation string) *ActionCardMessage {
	msg.ActionCard = ActionCardContent{
		Title:          title,
		Text:           text,
		Btns:           btns,
		BtnOrientation: btnOrientation,
	}
	return msg
}
