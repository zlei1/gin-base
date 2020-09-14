package routes

import (
	"github.com/gin-gonic/gin"

	"gin-base/pkg/websocket"
)

func InitWebsocket(router *gin.Engine) {
	go websocket.WsHub.Run()
	go websocket.WsHub.SendToClientService()
	go websocket.WsHub.SendToGroupService()
	go websocket.WsHub.SendToAllService()

	ws := router.Group("/ws")
	{
		ws.GET("/:channel", websocket.ClientConnect)
	}
}
