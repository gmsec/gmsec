package routers

import (
	"fmt"
	"net/http"
	"strings"

	"example/internal/api"
	"example/internal/service/hello"
	proto "example/rpc/example"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/gmsec/micro/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/xxjwxc/ginrpc"
	"github.com/xxjwxc/public/dev"
	"github.com/xxjwxc/public/tools"
)

// OnInitRoot 初始化
func OnInitRoot(s server.Server, router gin.IRoutes, objs ...interface{}) {
	var args []interface{}
	h := new(hello.Hello)
	args = append(args, h)
	proto.RegisterHelloServer(s, h) // 服务注册
	args = append(args, objs...)
	OnInitRouter(router, args...)
}

// OnInitRouter 默认初始化
func OnInitRouter(router gin.IRoutes, objs ...interface{}) {
	InitFunc(router)
	InitObj(router, objs...)
	init3rdGrpcHost()
}

// 初始化第三方host
func init3rdGrpcHost() {
	// micro.SetClientServiceAddr(haibuilder.GetHAIMotionBuilderName(), config.GetHaibuildHost()...)
	// micro.SetClientServiceName(haibuilder.GetHAIMotionBuilderName(), "haibuilder.srv.eg1")
}

// InitFunc 默认初始化函数
func InitFunc(router gin.IRoutes) {
	router.StaticFS("/file", http.Dir(tools.GetCurrentDirectory()+"/file")) //加载静态资源，一般是上传的资源，例如用户上传的图片
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	}) // 健康检查
	router.GET("/metrics", ginprom.PromHandler(promhttp.Handler())) // 添加grafana监控
}

// InitObj 初始化对象
func InitObj(router gin.IRoutes, objs ...interface{}) {
	base := ginrpc.New(ginrpc.WithCtx(api.NewAPIFunc), ginrpc.WithOutDoc(dev.IsDev()), ginrpc.WithDebug(dev.IsDev()), ginrpc.WithOutPath("internal/routers"),
		ginrpc.WithBeforeAfter(&ginrpc.DefaultGinBeforeAfter{}))

	base.OutDoc(true)
	base.Register(router, objs...) // 对象注册
}

// Cors 跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			//下面的都是乱添加的-_-~
			// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			// c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			// c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
