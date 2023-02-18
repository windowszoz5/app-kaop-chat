package main

import (
	"drone/common"
	"drone/compose"
	"drone/config"
	"drone/rpc"
	"drone/web/middleware"
	"drone/web/route"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	//初始环境配置
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
	conn, err := grpc.Dial("127.0.0.1:3366", grpc.WithInsecure(), grpc.WithStatsHandler(&common.ServerStats{}))
	if err != nil {
		fmt.Println("Dial err:14552454", err)
		return
	}
	defer conn.Close()
	common.ProductClient = rpc.NewProductClient(conn)

	// 配置路由
	route.Load(r)

	r.Run()
}
