package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"gin-base/pkg/utils"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Id    string
	Group string
	Conn  *websocket.Conn
	Send  chan []byte
}

func (c *Client) Read() {
	defer func() {
		WsHub.UnRegister <- c
		c.Conn.Close()
	}()

	for {
		messageType, message, err := c.Conn.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			log.Printf("ws client [%s] receive message error: %s", c.Id, err)
			break
		}

		log.Printf("ws client [%s] receive message: [%s]", c.Id, string(message))
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			log.Printf("ws client [%s] write message: [%s]", c.Id, string(message))

			err := c.Conn.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				log.Printf("ws client [%s] write message error: %s", c.Id, err)
			}
		}
	}
}

func ClientConnect(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &Client{
		Id:    utils.GenUuid(),
		Group: c.Param("channel"),
		Conn:  conn,
		Send:  make(chan []byte),
	}
	WsHub.Register <- client

	go client.Write()
	go client.Read()
}
