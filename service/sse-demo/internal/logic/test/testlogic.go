package test

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"sse-demo/service/sse-demo/internal/svc"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Test
func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test() (resp string, err error) {
	return "success", nil
}
