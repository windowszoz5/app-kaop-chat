package main

import (
	"drone/compose"
	"drone/config"
	"drone/route/middleware"
	"flag"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始环境配置
	var runConf string
	flag.StringVar(&runConf, "conf", "", "指定运行环境")
	flag.Parse()
	config.Init(runConf)

	//初始ES
	compose.InitEs()

	r := gin.Default()
	r.Use(middleware.MarkLog)

	// 配置路由
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": "1",
			"data": "data1 ",
		})
	})

	r.POST("/prd", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": "1",
			"data": "data1",
		})
	})

	r.Run()
}
