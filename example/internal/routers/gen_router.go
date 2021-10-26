package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1635251082)
	ginrpc.AddGenOne("Hello.SayHello", "hello.say_hello", []string{"post"})
}
