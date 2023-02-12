package route

import (
	"drone/web/handle"
	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine) {
	r.GET("/hello", handle.Hello)
}
