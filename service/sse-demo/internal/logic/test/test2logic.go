package test

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"sse-demo/service/sse-demo/internal/svc"
)

type Test2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Test2
func NewTest2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Test2Logic {
	return &Test2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Test2Logic) Test2() (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
