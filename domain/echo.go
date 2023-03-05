package domain

import (
	"drone/common"
	"drone/rpc"
)

func (q *dserver) Echo(req *HelloReq) *helloResp {
	common.ProductClient.Ping(q.Ctx, &rpc.Request{Ping: req.Hello})
	return nil
}
