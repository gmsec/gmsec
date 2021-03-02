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
	"github.com/xxjwxc/public/mydoc/myswagger"
	"github.com/xxjwxc/public/server"
)

// CallBack service call backe
func CallBack() {
	// swagger
	myswagger.SetHost("https://localhost:8080")
	myswagger.SetBasePath("example")
	myswagger.SetSchemes(true, false)
	// -----end --

	// reg := registry.NewDNSNamingRegistry()
	// grpc 相关 初始化服务
	service := micro.NewService(
		micro.WithName("example.srv.eg1"),
		// micro.WithRegisterTTL(time.Second*30),      //指定服务注册时间
		micro.WithRegisterInterval(time.Second*15), //让服务在指定时间内重新注册
		// micro.WithRegistryNaming(reg),
	)
	// ----------- end

	// gin restful 相关
	router := gin.Default()
	router.Use(routers.Cors())
	v1 := router.Group("/example/api/v1")
	routers.OnInitRoot(service.Server(), v1) // 自定义初始化
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
