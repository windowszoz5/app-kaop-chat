package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func CheckWebSocket(ctx *gin.Context) {
	if ctx.IsWebsocket() == false {
		ctx.Abort()
		return
	}

	//校验来源
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			fmt.Println("校验来源")
			return true
		},
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.Abort()
		return
	}
	ctx.Set("conn", conn)
}
