package domain

import "context"

type Realization interface {
	Hello(req *HelloReq) *helloResp
}

type dserver struct {
	Ctx context.Context
}

func GetDserver(ctx context.Context) *dserver {
	return &dserver{
		Ctx: ctx,
	}
}
