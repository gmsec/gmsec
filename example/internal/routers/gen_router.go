package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1614677081)
	ginrpc.AddGenOne("Hello.SayHello", "hello.say_hello", []string{"post"})
}
