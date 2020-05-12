package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/xxjwxc/public/dev"

	_ "gmsec/internal/routers" // debug模式需要添加[mod]/routers 注册注解路由

	"context"

	proto "gmsec/rpc"

	"github.com/gin-gonic/gin"
	"github.com/gmsec/goplugins/plugin"
	"github.com/gmsec/micro"
	"github.com/xxjwxc/ginrpc"
	"github.com/xxjwxc/gowp/workpool"
	"github.com/xxjwxc/public/mydoc/myswagger"
)

// TestMain test main
func TestMain(t *testing.T) {
	// swagger
	myswagger.SetHost("https://localhost:8080")
	myswagger.SetBasePath("gmsec")
	myswagger.SetSchemes(true, false)
	// -----end --

	// reg := registry.NewDNSNamingRegistry()
	// grpc 相关 初始化服务
	service := micro.NewService(
		micro.WithName("lp.srv.eg1"),
		micro.WithRegisterTTL(time.Second*30), //指定服务注册时间
		// micro.WithRegisterInterval(time.Second*15), //让服务在指定时间内重新注册
		// micro.WithRegistryNameing(reg),
	)
	h := new(hello)
	proto.RegisterHelloServer(service.Server(), h) // 服务注册
	// ----------- end

	// gin restful 相关
	base := ginrpc.New(ginrpc.WithCtx(func(c *gin.Context) interface{} {
		return context.Background()
	}), ginrpc.WithDebug(dev.IsDev()), ginrpc.WithGroup("xxjwxc"))
	router := gin.Default()
	base.Register(router, h) // 对象注册
	// ------ end

	plg, _ := plugin.Run(plugin.WithMicro(service),
		plugin.WithGin(router),
		plugin.WithAddr("localhost:8080"))

	clientTest() // client test

	time.Sleep(3 * time.Second)
	plg.Stop()
	fmt.Println("done")
}

func clientTest() {
	// first
	// service := micro.NewService(
	// 	micro.WithName("lp.srv.eg1"),
	// 	// micro.WithRegisterTTL(time.Second*30),      //指定服务注册时间
	// 	micro.WithRegisterInterval(time.Second*15), //让服务在指定时间内重新注册
	// 	//micro.WithRegistryNameing(reg),
	// )
	go func() {
		wp := workpool.New(2)     //设置最大线程数
		for i := 0; i < 20; i++ { //开启20个请求
			wp.Do(func() error {
				run()
				return nil
			})
		}

		wp.Wait()
		fmt.Println("clinet down")
	}()
}

func run() {
	micro.SetClientServiceName(proto.GetHelloName(), "lp.srv.eg1") // set client group
	say := proto.GetHelloClient()

	var request proto.HelloRequest
	r := rand.Intn(500)
	request.Name = fmt.Sprintf("%v", r)

	ctx := context.Background()

	for i := 0; i < 10; i++ {
		resp, err := say.SayHello(ctx, &request)
		if err != nil {
			fmt.Println("==========err:", err)
		}
		fmt.Println(resp)
		time.Sleep(1 * time.Second)
	}
}
