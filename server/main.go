package main

import (
	"context"
	"drone/common"
	"drone/server/rpc"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type Product struct {
	rpc.UnimplementedProductServer
}

func (Product) Ping(ctx context.Context, req *rpc.Request) (*rpc.Response, error) {
	fmt.Println(ctx.Value(common.X_REQ_UUID), "服务端接收到")
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 3366))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 实例化server
	s := grpc.NewServer()
	rpc.RegisterProductServer(s, &Product{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
