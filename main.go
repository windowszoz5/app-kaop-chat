package main

import (
	"context"
	"drone/common"
	"drone/compose"
	"drone/config"
	"drone/route/middleware"
	"drone/server/rpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
)

func main() {
	//初始环境配置
	var runConf string = "./config/master.json"
	//flag.StringVar(&runConf, "conf", "", "指定运行环境")
	//flag.Parse()
	config.Init(runConf)

	//初始ES
	compose.InitEs()

	//中间件
	r := gin.Default()
	r.Use(middleware.MarkLog)

	//连接微服务服务端
	conn, err := grpc.Dial("127.0.0.1:3366", grpc.WithInsecure(), grpc.WithStatsHandler(&serverStats{}))
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	productClient := rpc.NewProductClient(conn)

	// 配置路由
	r.GET("/hello", func(c *gin.Context) {
		_, err := productClient.Ping(c, &rpc.Request{Ping: "1"})
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

type serverStats struct{}

func (h *serverStats) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	md := metadata.New(map[string]string{common.X_REQ_UUID: fmt.Sprintf("%v", ctx.Value(common.X_REQ_UUID))})
	return metadata.NewOutgoingContext(ctx, md)
}

func (h *serverStats) HandleRPC(ctx context.Context, s stats.RPCStats) {

}

func (h *serverStats) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	return ctx
}

func (h *serverStats) HandleConn(ctx context.Context, s stats.ConnStats) {
}
