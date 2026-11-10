package hello

import (
	"context"
	proto "example/rpc/example"

	"github.com/gmsec/micro/tracer"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"

	"google.golang.org/grpc/metadata"
)

type Hello struct {
}

// SayHello 链路追踪示例: 进来自带的 ctx 已有服务端 span, 往下游扔就是把这个 ctx 传下去
func (h *Hello) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx) // 只有 gRPC 进来的请求才有上游 metadata(含 trace 头)
	mylog.Debugf("SayHello metadata=%v present=%v req=%v", md, ok, req)

	newCtx, finish, err := tracer.Start("SayHello", ctx) // 开业务子 span
	if err != nil {
		mylog.Error(err)
	}
	if newCtx != nil { // 保留新上下文, 下游调用作为业务 span 的子 span
		ctx = newCtx
	}
	defer func() {
		if finish != nil {
			finish()
		}
	}()

	// 下游 gRPC: 把 ctx 传进去就行(micro 的 client 拦截器会自动把 span 注入 metadata)
	// resp, err := proto.GetExampleClient().SayHello(ctx, &proto.HelloRequest{Name: req.Name})

	// 下游 HTTP: 不会自动注入, 要手动把 span 塞进 header, 见 internal/api/user.go 的 postJSONWithTrace

	return &proto.HelloReply{
		Message: tools.GetRandomString(8),
	}, nil
}
