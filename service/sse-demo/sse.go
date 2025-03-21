package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"sse-demo/service/sse-demo/internal/config"
	"sse-demo/service/sse-demo/internal/handler"
	"sse-demo/service/sse-demo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/sse-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithFileServer("/static", http.Dir("static")))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 初始化 SSE 处理
	sseHandler := NewSseHandler()

	// 注册 SSE 路由
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/sse",
		Handler: sseHandler.Serve,
	}, rest.WithTimeout(0))

	// 在单独的 goroutine 中模拟事件
	go sseHandler.SimulateEvents()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

type SseHandler struct {
	clients map[chan string]struct{}
}

func NewSseHandler() *SseHandler {
	return &SseHandler{
		clients: make(map[chan string]struct{}),
	}
}

// Serve 处理 SSE 连接
func (h *SseHandler) Serve(w http.ResponseWriter, r *http.Request) {
	// 设置 SSE 必需的 HTTP 头
	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Connection", "keep-alive")

	// 为每个客户端创建一个 channel
	clientChan := make(chan string)
	h.clients[clientChan] = struct{}{}

	// 客户端断开时清理
	defer func() {
		delete(h.clients, clientChan)
		close(clientChan)
	}()

	// 持续监听并推送事件
	for {
		select {
		case msg := <-clientChan:
			// 发送事件数据
			fmt.Fprintf(w, "data: %s\n\n", msg)
			w.(http.Flusher).Flush()
		case <-r.Context().Done():
			// 客户端断开连接
			return
		}
	}
}

// SimulateEvents 模拟周期性事件
func (h *SseHandler) SimulateEvents() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		message := fmt.Sprintf("Server time: %s", time.Now().Format(time.RFC3339))
		// 广播给所有客户端
		for clientChan := range h.clients {
			select {
			case clientChan <- message:
			default:
				// 跳过阻塞的 channel
			}
		}
	}
}
