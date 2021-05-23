package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1621758764)
	ginrpc.AddGenOne("Hello.SayHello", "hello.say_hello", []string{"post"})
}
