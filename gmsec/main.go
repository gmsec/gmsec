package main

import (
	"fmt"
	"os"
	"time"

	proto "github.com/gmsec/gmsec/rpc"

	"github.com/xxjwxc/public/dev"

	"gmsec/internal/config"
	_ "gmsec/internal/routers" // debug模式需要添加[mod]/routers 注册注解路由

	"context"

	"github.com/gin-gonic/gin"
	"github.com/gmsec/goplugins/plugin"
	"github.com/gmsec/micro"
	"github.com/xxjwxc/ginrpc"
	"github.com/xxjwxc/public/mydoc/myswagger"
	"github.com/xxjwxc/public/server"
)

// CallBack service call backe
func CallBack() {
	// swagger
	myswagger.SetHost("https://localhost:8080")
	myswagger.SetBasePath("gmsec")
	myswagger.SetSchemes(true, false)
	// -----end --

	// reg := registry.NewDNSNamingRegistry()
	// grpc 相关 初始化服务
	service := micro.NewService(
		micro.WithName("lp.srv.eg1"),
		// micro.WithRegisterTTL(time.Second*30),      //指定服务注册时间
		micro.WithRegisterInterval(time.Second*15), //让服务在指定时间内重新注册
		//micro.WithRegistryNameing(reg),
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

	plg, b := plugin.Run(plugin.WithMicro(service),
		plugin.WithGin(router),
		plugin.WithAddr(":82"))

	if b == nil {
		plg.Wait()
	}
	fmt.Println("done")
}

func main() {
	if config.GetIsDev() || len(os.Args) == 0 {
		CallBack()
	} else {
		server.On(config.GetServiceConfig()).Start(CallBack)
	}
}
