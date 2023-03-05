package route

import (
	"drone/web/handle"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"
)

func Load(r *gin.Engine) {
	r.GET("/hello", handle.Hello)
	r.GET("/echo", echo2)
}

func echo2(c *gin.Context) {
	//升级get请求为webSocket协议
	conn, _ := c.Get("conn")
	fmt.Println("类型")
	ws, _ := conn.(*websocket.Conn)
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			c.Writer.Write([]byte(err.Error()))
			break
		}
		fmt.Println("client message " + string(message))
		//写入ws数据
		err = ws.WriteMessage(mt, []byte(time.Now().String()))
		if err != nil {
			break
		}
		fmt.Println("system message " + time.Now().String())
	}
}
