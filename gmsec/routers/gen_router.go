package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1589276483)
	ginrpc.AddGenOne("hello.SayHello", "hello.say_hello", []string{"post"})
}
