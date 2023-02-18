package common

import (
	"context"
	"drone/rpc"
	"fmt"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
)

type ServerStats struct{}

var ProductClient rpc.ProductClient

func (h *ServerStats) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	md := metadata.New(map[string]string{X_TRACK_ID: fmt.Sprint(ctx.Value(X_TRACK_ID))})
	outCtx := metadata.NewOutgoingContext(ctx, md)
	return outCtx
}

func (h *ServerStats) HandleRPC(ctx context.Context, s stats.RPCStats) {

}

func (h *ServerStats) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	return ctx
}

func (h *ServerStats) HandleConn(ctx context.Context, s stats.ConnStats) {
}
