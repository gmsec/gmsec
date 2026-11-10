package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example/internal/config"
	"example/internal/core"
	"example/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/gmsec/goplugins/plugin"
	"github.com/gmsec/micro"
	"github.com/gmsec/micro/registry"
	"github.com/gmsec/micro/tracer"
	"github.com/xxjwxc/public/dev"
	"github.com/xxjwxc/public/mydoc/myswagger"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/server"
)

// CallBack service call backe
func CallBack() {
	// swagger
	myswagger.SetHost("https://localhost:" + config.GetPort())
	myswagger.SetBasePath("example")
	myswagger.SetSchemes(true, false)
	// -----end --
	tracer.WithTracer(config.GetJaeger().Tag, config.GetJaeger().Addr, config.GetJaeger().Percent, config.GetJaeger().MaxTagsLen) // 链路追踪

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
	trace := tracer.GetTracer()
	if trace != nil { // 链路追踪
		router.Use(routers.UseJager(trace))
	}
	v1 := router.Group("/example/api/v1")
	routers.OnInitRoot(service.Server(), v1) // 自定义初始化
	// ------ end

	plg, err := plugin.Run(plugin.WithMicro(service),
		plugin.WithGin(router),
		plugin.WithAddr(":"+config.GetPort()))
	if err != nil {
		// 启动失败必须记日志并以非 0 退出;历史实现 fmt.Println 后直接 return,
		// 服务根本没起来却当成正常结束(前台模式会直接 exit 0)。
		mylog.Error(err)
		core.Dao.Destroy()
		tracer.CloseTracer()
		os.Exit(1)
	}

	// 捕获 SIGINT/SIGTERM,收到信号后主动触发优雅退出。
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	done := make(chan struct{})
	go func() {
		plg.Wait()
		close(done)
	}()

	select {
	case <-done:
		mylog.Info("服务已退出,开始清理")
	case <-ctx.Done():
		mylog.Info("收到退出信号,开始优雅关闭")
	}

	// 退出顺序:先停服务摘流量,再释放数据库连接,最后关闭 tracer。
	plg.Stop()
	core.Dao.Destroy()
	tracer.CloseTracer()
	mylog.Info("done")
}

func main() {
	// 配置加载错误在包 init 中保存,由 main 决定退出码(非 0)。
	if err := config.ConfigError(); err != nil {
		mylog.Error(err)
		os.Exit(1)
	}

	// len(os.Args) == 0 恒为假(os.Args 至少含程序名),原来那个条件永远不生效。
	// 改用显式的 debug 参数控制前台模式,和 Makefile 里的 `./example debug` 对齐。
	debug := flag.Bool("debug", false, "前台运行(不走系统服务管理)")
	flag.Parse()
	foreground := *debug
	for _, a := range flag.Args() {
		if a == "debug" {
			foreground = true
		}
	}

	if config.GetIsDev() || foreground {
		// 与历史 `./example debug` 行为一致:前台调试时打开 dev 标记(swagger/ginrpc 调试输出)。
		if foreground {
			dev.OnSetDev(true)
		}
		CallBack()
		return
	}
	server.On(config.GetServiceConfig()).Start(CallBack)
}
