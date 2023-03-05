package handle

import (
	"drone/common"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Echo(c *gin.Context) {
	//升级get请求为webSocket协议
	conn, _ := c.Get("conn")
	ws, _ := conn.(*websocket.Conn)
	common.SetConn("list", ws, 1)
}
