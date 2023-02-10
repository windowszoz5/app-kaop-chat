package main

import (
	"drone/compose"
	"drone/config"
	"drone/route/middleware"
	"drone/server/rpc"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// 初始环境配置
	var runConf string
	flag.StringVar(&runConf, "conf", "", "指定运行环境")
	flag.Parse()
	config.Init(runConf)

	//初始ES
	compose.InitEs()

	//中间件
	r := gin.Default()
	r.Use(middleware.MarkLog)

	//连接微服务服务端
	conn, err := grpc.Dial("localhost:3366", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()
	productClient := rpc.NewProductClient(conn)

	// 配置路由
	r.GET("/hello", func(c *gin.Context) {
		_, err := productClient.Ping(c.Request.Context(), &rpc.Request{Ping: "1"})
		if err != nil {
			return
		}
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
