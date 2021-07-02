package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1625219438)
	ginrpc.AddGenOne("Hello.SayHello", "hello.say_hello", []string{"post"})
}
