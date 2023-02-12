package handle

import (
	"drone/common"
	"drone/server/rpc"
	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	common.ProductClient.Ping(ctx, &rpc.Request{Ping: "666"})
	ctx.JSON(200, map[string]string{
		"helo": "1",
	})
}
