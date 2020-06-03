package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/xxjwxc/gowp/workpool"
	"github.com/xxjwxc/public/dev"

	"context"

	proto "gmsec/rpc/example"

	"github.com/gin-gonic/gin"
	"github.com/gmsec/goplugins/api"
	"github.com/gmsec/goplugins/plugin"
	"github.com/gmsec/micro"
	"github.com/xxjwxc/ginrpc"
	"github.com/xxjwxc/public/mydoc/myswagger"
)

// TestMain test main
func TestServer(m *testing.T) {
	// swagger
	myswagger.SetHost("https://localhost:8080")
	myswagger.SetBasePath("gmsec")
	myswagger.SetSchemes(true, false)
	// -----end --

	// reg := registry.NewDNSNamingRegistry()
	// grpc 相关 初始化服务
	service := micro.NewService(
		micro.WithName("xxjwxc.lp.srv.eg1"),
		micro.WithRegisterTTL(time.Second*30), //指定服务注册时间
		// micro.WithRegisterInterval(time.Second*15), //让服务在指定时间内重新注册
		// micro.WithRegistryNameing(reg),
	)
	h := new(hello)
	proto.RegisterHelloServer(service.Server(), h) // 服务注册
	// ----------- end

	// gin restful 相关
	base := ginrpc.New(ginrpc.WithCtx(api.NewAPIFunc), ginrpc.WithDebug(dev.IsDev()))
	router := gin.Default()
	v1 := router.Group("/xxjwxc/api/v1")
	base.Register(v1, h) // 对象注册
	// ------ end

	plg, _ := plugin.Run(plugin.WithMicro(service),
		plugin.WithGin(router),
		plugin.WithAddr("localhost:8080"))
	defer plg.Stop()
	// TestClient() // client test
	plg.Wait()
	/*time.Sleep(3 * time.Second)
	plg.Stop()
	fmt.Println("done")*/
}

func TestClient(m *testing.T) {
	micro.SetClientServiceName(proto.GetHelloName(), "lp.srv.eg1") // set client group
	// first
	micro.NewService(
		micro.WithName("lp.srv.eg1"),
		// micro.WithRegisterTTL(time.Second*30),      //指定服务注册时间
		micro.WithRegisterInterval(time.Second*15), //让服务在指定时间内重新注册
		//micro.WithRegistryNameing(reg),
	)
	wp := workpool.New(2)     //设置最大线程数
	for i := 0; i < 20; i++ { //开启20个请求
		wp.Do(func() error {
			say := proto.GetHelloClient()
			var request proto.HelloRequest
			request.Name = fmt.Sprintf("%v", rand.Intn(500))

			ctx := context.Background()
			resp, err := say.SayHello(ctx, &request)
			if err != nil {
				fmt.Println("==========err:", err)
			}
			fmt.Println(resp)
			return nil
		})
	}
	wp.Wait()
}

func run() {
	say := proto.GetHelloClient()
	var request proto.HelloRequest
	request.Name = fmt.Sprintf("%v", rand.Intn(500))

	ctx := context.Background()
	resp, err := say.SayHello(ctx, &request)
	if err != nil {
		fmt.Println("==========err:", err)
	}
	fmt.Println(resp)
	time.Sleep(1 * time.Second)
}
