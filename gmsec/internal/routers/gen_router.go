package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1589185123)
	ginrpc.AddGenOne("Hello.Block", "/block", []string{"post", "get"})
}
