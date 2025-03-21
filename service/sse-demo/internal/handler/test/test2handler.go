package test

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sse-demo/service/sse-demo/internal/logic/test"
	"sse-demo/service/sse-demo/internal/svc"
)

// Test2
func Test2Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := test.NewTest2Logic(r.Context(), svcCtx)
		resp, err := l.Test2()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
