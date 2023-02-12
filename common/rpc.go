package common

import (
	"context"
	"drone/server/rpc"
	"fmt"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
)

type ServerStats struct{}

var ProductClient rpc.ProductClient

func (h *ServerStats) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	md := metadata.New(map[string]string{X_REQ_UUID: fmt.Sprintf("%v", ctx.Value(X_REQ_UUID))})
	return metadata.NewOutgoingContext(ctx, md)
}

func (h *ServerStats) HandleRPC(ctx context.Context, s stats.RPCStats) {

}

func (h *ServerStats) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	return ctx
}

func (h *ServerStats) HandleConn(ctx context.Context, s stats.ConnStats) {
}
