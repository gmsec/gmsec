package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1614679563)
	ginrpc.AddGenOne("Hello.SayHello", "hello.say_hello", []string{"post"})
}
