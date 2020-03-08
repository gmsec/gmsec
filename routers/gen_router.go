package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1583685971)
	ginrpc.AddGenOne("Hello.Block", "/block", []string{"post", "get"})
	ginrpc.AddGenOne("Hello.MySql2", "hello.my_sql2", []string{"post"})
}
