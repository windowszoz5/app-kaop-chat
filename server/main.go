package main

import (
	"context"
	"drone/server/rpc"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type Product struct {
	rpc.UnimplementedProductServer
}

func (Product) Ping(ctx context.Context, req *rpc.Request) (*rpc.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func main() {
	// 监听端口
	url := fmt.Sprintf("localhost:%d", 3366)
	lis, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 实例化server
	s := grpc.NewServer(grpc.StatsHandler(&serverStats{}))

	rpc.RegisterProductServer(s, &Product{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type serverStats struct{}

func (h *serverStats) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	fmt.Println("333333")
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println("值", md)
	return ctx
}

func (h *serverStats) HandleRPC(ctx context.Context, s stats.RPCStats) {
	//fmt.Println(22222)
}

func (h *serverStats) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	//fmt.Println("11111111111")
	return context.TODO()
}

func (h *serverStats) HandleConn(ctx context.Context, s stats.ConnStats) {
	//fmt.Println("?????????????????") // Returns nil, can't access the value

	switch s.(type) {
	case *stats.ConnEnd:
		fmt.Println("client disconnected")
		break
	}
}
