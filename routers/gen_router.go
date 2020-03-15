package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1584261661)
	ginrpc.AddGenOne("Hello.Block", "/block", []string{"post", "get"})
}
