package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1589262558)
	ginrpc.AddGenOne("hello.SayHello", "hello.say_hello", []string{"post"})
}
