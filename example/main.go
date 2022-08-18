package main

import (
	"fmt"
	"os"
	"time"

	"example/internal/config"
	"example/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/gmsec/goplugins/plugin"
	"github.com/gmsec/micro"
	"github.com/gmsec/micro/registry"
	"github.com/xxjwxc/public/mydoc/myswagger"
	"github.com/xxjwxc/public/server"
	_ "go.etcd.io/etcd/client/v3"
	// "github.com/gmsec/micro/tracer"
)

// CallBack service call backe
func CallBack() {
	// swagger
	myswagger.SetHost("https://localhost:" + config.GetPort())
	myswagger.SetBasePath("example")
	myswagger.SetSchemes(true, false)
	// -----end --
	// tracer.WithTracer("192.155.1.150:6831") // 链路追踪
	reg := registry.NewDNSNamingRegistry()
	// reg := etcdv3.NewEtcdv3NamingRegistry(v3.Config{
	// 	Endpoints:   config.GetEtcdInfo().Addrs,
	// 	DialTimeout: time.Second * time.Duration(config.GetEtcdInfo().Timeout),
	// })
	// grpc 相关 初始化服务
	service := micro.NewService(
		micro.WithName("gmsec.srv.example"),
		// micro.WithRegisterTTL(time.Second*30),      //指定服务注册时间
		micro.WithRegisterInterval(time.Second*15), //让服务在指定时间内重新注册
		micro.WithRegistryNaming(reg),
	)
	// ----------- end

	// gin restful 相关
	router := gin.Default()
	router.Use(routers.Cors())
	// trace := tracer.GetTracer()
	// if trace != nil { // 链路追踪
	// 	router.Use(routers.UseJager(trace))
	// }
	v1 := router.Group("/example/api/v1")
	routers.OnInitRoot(service.Server(), v1) // 自定义初始化
	// ------ end

	plg, b := plugin.Run(plugin.WithMicro(service),
		plugin.WithGin(router),
		plugin.WithAddr(":"+config.GetPort()))

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
