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

	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			fmt.Println("校验来源")
			return true
		},
	}
	conn, _ := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	ctx.Set("conn", conn)

}
