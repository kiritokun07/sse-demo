## sse-demo

#### 搞定 Go 实时消息推送
https://mp.weixin.qq.com/s/tf5Opz2I8VCu5n24YjXUDw

http://localhost:8888/static/

```shell
goctl api go -api sse-demo.api -dir ./service/sse-demo
```

```markdown
# 安装go
https://go.dev/

# 设置七牛云源
go env -w GOPROXY=https://goproxy.cn,direct
# 下载好git项目后，在go.mod所在目录下
go mod tidy
# 运行go
go run .\main.go
# 测试go
go test hello_test.go

# 安装goctl
go install github.com/zeromicro/go-zero/tools/goctl@latest
# 高版本的goctl需要关闭实验功能，否则执行很慢而且int64变成uint64
goctl env -w GOCTL_EXPERIMENTAL=off
```