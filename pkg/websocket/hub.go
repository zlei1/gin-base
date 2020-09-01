package websocket

import (
	"log"
	"sync"
)

var WsHub = NewHub()

type Hub struct {
	Lock             sync.Mutex
	Group            map[string]map[string]*Client
	GroupCount       uint
	ClientCount      uint
	Register         chan *Client
	UnRegister       chan *Client
	Message          chan *MessageData
	GroupMessage     chan *GroupMessageData
	BroadcastMessage chan *BroadcastMessageData
}

type MessageData struct {
	Id      string
	Group   string
	Message []byte
}

type GroupMessageData struct {
	Group   string
	Message []byte
}

type BroadcastMessageData struct {
	Message []byte
}

func NewHub() *Hub {
	return &Hub{
		Group:            make(map[string]map[string]*Client),
		GroupCount:       0,
		ClientCount:      0,
		Register:         make(chan *Client),
		UnRegister:       make(chan *Client),
		Message:          make(chan *MessageData),
		GroupMessage:     make(chan *GroupMessageData),
		BroadcastMessage: make(chan *BroadcastMessageData),
	}
}

func (h *Hub) Run() {
	for {
		select {
		// 注册
		case client := <-h.Register:
			log.Printf("ws register client [%s] to group [%s]", client.Id, client.Group)

			h.Lock.Lock()
			if h.Group[client.Group] == nil {
				h.Group[client.Group] = make(map[string]*Client)
				h.GroupCount += 1
			}
			h.Group[client.Group][client.Id] = client
			h.ClientCount += 1
			h.Lock.Unlock()
		// 注销
		case client := <-h.UnRegister:
			log.Printf("ws unregister client [%s] from group [%s]", client.Id, client.Group)

			h.Lock.Lock()
			if _, ok := h.Group[client.Group]; ok {
				if _, ok := h.Group[client.Group][client.Id]; ok {
					close(client.Send)
					delete(h.Group[client.Group], client.Id)
					h.ClientCount -= 1
					if len(h.Group[client.Group]) == 0 {
						delete(h.Group, client.Group)
						h.GroupCount -= 1
					}
				}
			}
			h.Lock.Unlock()
		}
	}
}
