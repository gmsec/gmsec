[![Build Status](https://travis-ci.org/xxjwxc/gmsec.svg?branch=master)](https://travis-ci.org/xxjwxc/gmsec)
[![Go Report Card](https://goreportcard.com/badge/github.com/xxjwxc/gmsec)](https://goreportcard.com/report/github.com/xxjwxc/gmsec)
[![codecov](https://codecov.io/gh/xxjwxc/gmsec/branch/master/graph/badge.svg)](https://codecov.io/gh/xxjwxc/gmsec)
[![GoDoc](https://godoc.org/github.com/xxjwxc/gmsec?status.svg)](https://godoc.org/github.com/xxjwxc/gmsec)


# [gmsec](https://github.com/gmsec/gmsec)


### golang 微服务集成框架 

- [grpc](https://github.com/grpc/grpc-go)
- [gorm 自动构建](https://github.com/xxjwxc/gormt)
- [grpc 参数自动绑定工具](https://github.com/xxjwxc/ginrpc)
- [dns 注册发现](https://github.com/asim/mdns)
- [markdown/mindoc 文档自动导出](https://github.com/grpc)


## 安装

- install

- proto环境安装

```
 make install 
```

- 本地环境搭建(gmsec为例)

```
make gen
```

## proto定义

```
syntax = "proto3"; // 指定proto版本
package proto;     // 指定包名
option go_package = ".;proto"; // 指定路径

// 定义Hello服务
service Hello {
    // 定义SayHello方法
    rpc SayHello(HelloRequest) returns (HelloReply) {}
}
// HelloRequest 请求结构
message HelloRequest {
    string name = 1; // 名字
}
// HelloReply 响应结构
message HelloReply {
    string message = 1; // 消息
}
```

## 服务端代码示例

```
package main

import (
	"context"
	"fmt"
	proto "gmsec/rpc"

	"github.com/gin-gonic/gin"
	"github.com/gmsec/goplugins/plugin"
	"github.com/gmsec/micro"
	"github.com/xxjwxc/ginrpc"
)

func main() {
	// reg := registry.NewDNSNamingRegistry()
	// grpc 相关 初始化服务
	service := micro.NewService(
		micro.WithName("lp.srv.eg1"),
		// micro.WithRegisterTTL(time.Second*30),      //指定服务注册时间
		// micro.WithRegisterInterval(time.Second*15), //让服务在指定时间内重新注册
		// micro.WithRegistryNameing(reg),
	)
	h := new(hello)
	proto.RegisterHelloServer(service.Server(), h) // 服务注册
	// ----------- end

	// gin restful 相关
	base := ginrpc.New(ginrpc.WithCtx(Ctx), ginrpc.WithDebug(true), ginrpc.WithGroup("xxjwxc"))
	router := gin.Default()  // gen 对象
	base.Register(router, h) // genrpc对象注册
	// ------ end

	plg, b := plugin.Run(plugin.WithMicro(service),
		plugin.WithGin(router),
		plugin.WithAddr(":8080")) // 开始服务(公用端口)

	if b == nil {
		plg.Wait()
	}
	fmt.Println("done")
}

func Ctx(c *gin.Context) interface{} {
	return context.Background()
}
```

## 客户端代码:

```
package main

import (
	"context"
	"fmt"
	proto "gmsec/rpc"

	"github.com/gmsec/micro"
	// "github.com/gmsec/micro/registry"
)

func main() {
	// reg := registry.NewDNSNamingRegistry()
	// grpc 相关 注册服务发现等
	micro.NewService(
		micro.WithName("lp.srv.eg1"),
		// micro.WithRegisterTTL(time.Second*30),      //指定服务注册时间
		// micro.WithRegisterInterval(time.Second*15), //让服务在指定时间内重新注册
		// micro.WithRegistryNameing(reg),
	)
	// ----------- end

	micro.SetClientServiceName(proto.GetHelloName(), "lp.srv.eg1") // set client group
	say := proto.GetHelloClient()

	var request proto.HelloRequest
	request.Name = "xxjwxc"
	ctx := context.Background()
	resp, _ := say.SayHello(ctx, &request)
	fmt.Println("result:", resp)
}
```

## 更多示例 => [传送门](https://github.com/gmsec/gmsec/tree/master/gmsec)

## 正在做
- etcdv3

- ###### [传送门](https://github.com/gmsec)


## 欢迎一起共建共享