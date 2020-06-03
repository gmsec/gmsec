package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1591189220)
	ginrpc.AddGenOne("hello.SayHello", "hello.say_hello", []string{"post"})
}
