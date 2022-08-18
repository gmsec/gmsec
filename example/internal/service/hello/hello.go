package hello

import (
	"context"
	proto "example/rpc/example"
	"fmt"

	"github.com/xxjwxc/public/tools"

	"google.golang.org/grpc/metadata"
)

type Hello struct {
}

func (h *Hello) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	// c := api.WidthContext(ctx)
	// tp, _ := c.GetGinCtx().Get("jager_ctx")
	// ctx = tp.(context.Context)
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	fmt.Println(ok)
	fmt.Println(req)
	return &proto.HelloReply{
		Message: tools.GetRandomString(8),
	}, nil
}
