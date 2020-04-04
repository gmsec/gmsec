package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1586003407)
	ginrpc.AddGenOne("Hello.Block", "/block", []string{"post", "get"})
}
