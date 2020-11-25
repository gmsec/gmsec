package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1606289928)
	ginrpc.AddGenOne("Hello.SayHello", "hello.say_hello", []string{"post"})
}
