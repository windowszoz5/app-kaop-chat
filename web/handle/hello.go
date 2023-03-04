package handle

import (
	"drone/domain"
	//cjh "gitee.com/poppin-jch/cjh"
	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	hello := ctx.GetString("hello")
	domain.GetDserver(ctx).Hello(&domain.HelloReq{
		Hello: hello,
	})
	//ctx.JSON(200, cjh.Json().ReplaceData(data))
}
