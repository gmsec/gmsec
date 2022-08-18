
	package routers
	
	import (
		"github.com/xxjwxc/ginrpc"
	)
	
	func init() {
		ginrpc.SetVersion(1660818462)
		ginrpc.AddGenOne("Hello.SayHello", "hello.say_hello", []string{ "post" })
		 }
	