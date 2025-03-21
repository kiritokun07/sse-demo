package test

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sse-demo/service/sse-demo/internal/logic/test"
	"sse-demo/service/sse-demo/internal/svc"
)

// Test
func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := test.NewTestLogic(r.Context(), svcCtx)
		resp, err := l.Test()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
