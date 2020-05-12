package main

import (
	"context"
	"fmt"
	proto "gmsec/rpc"

	"github.com/xxjwxc/public/tools"

	"google.golang.org/grpc/metadata"
)

type hello struct {
}

func (h *hello) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	fmt.Println(ok)
	fmt.Println(ctx.Value("HELLO"))
	fmt.Println(ctx.Value("WROLD"))
	fmt.Println(req)
	return &proto.HelloReply{
		Message: tools.GetRandomString(8),
	}, nil
}
