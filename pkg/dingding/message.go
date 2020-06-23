package dingding

const (
	MsgTypeText       string = "text"
	MsgTypeMarkdown   string = "markdown"
	MsgTypeLink       string = "link"
	MsgTypeActionCard string = "actionCard"
	MsgTypeFeedCard   string = "feedCard"
)

type Message interface {
	ToJsonByte() ([]byte, error)
}

type At struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}
