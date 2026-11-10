package routers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"example/internal/api"
	"example/internal/service/hello"
	proto "example/rpc/example"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/gmsec/micro/server"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/xxjwxc/ginrpc"
	"github.com/xxjwxc/public/dev"
	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	"google.golang.org/grpc/status"
)

// OnInitRoot 初始化
func OnInitRoot(s server.Server, router gin.IRoutes, objs ...interface{}) {
	var args []interface{}
	h := new(hello.Hello)
	args = append(args, h)
	proto.RegisterExampleServer(s, h) // 服务注册
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
	// micro.SetClientServiceName(oauth2.GetOauth2Name(), "gmsec.srv.oauth2")
}

// InitFunc 默认初始化函数
func InitFunc(router gin.IRoutes) {
	// /file 原来无条件挂载一个并不存在的目录,请求必然 404,还把目录约定暴露出去。
	// 只有目录真实存在时才挂载,并挂上 token 鉴权中间件,避免静态资源被匿名读取。
	fileDir := filepath.Join(tools.GetCurrentDirectory(), "file")
	if fi, err := os.Stat(fileDir); err == nil && fi.IsDir() {
		router.Use(fileAuth())
		router.StaticFS("/file", gin.Dir(fileDir, false))
	}
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	}) // 健康检查
	router.GET("/metrics", ginprom.PromHandler(promhttp.Handler())) // 添加grafana监控
}

// fileAuth 静态资源鉴权:只对 /file 下的请求生效,复用业务 token 校验。
func fileAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.URL.Path, "/file") {
			c.Next()
			return
		}
		token, err := c.Cookie(api.UserToken)
		if err != nil {
			token, err = c.Cookie(api.SessionToken)
		}
		if err != nil || token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if _, ok := api.GetUserFromToken(api.TraceCtx(c), token); !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

// InitObj 初始化对象
func InitObj(router gin.IRoutes, objs ...interface{}) {
	base := ginrpc.New(ginrpc.WithCtx(api.NewAPIFunc), ginrpc.WithOutDoc(dev.IsDev()), ginrpc.WithDebug(dev.IsDev()),
		ginrpc.WithOutPath("internal/routers"), ginrpc.WithImportFile("rpc/common", "../apidoc/rpc/common"),
		ginrpc.WithBeforeAfter(&JagerGinBeforeAfter{}))

	base.OutDoc(true)
	base.Register(router, objs...) // 对象注册
}

// allowedOrigins 跨域来源白名单,来自环境变量 GMSEC_CORS_ALLOWED_ORIGINS(逗号分隔)。
// 不配置时为空,此时不回显 Access-Control-Allow-Origin,等价于只允许同源。
func allowedOrigins() []string {
	var out []string
	for _, s := range strings.Split(os.Getenv("GMSEC_CORS_ALLOWED_ORIGINS"), ",") {
		if s = strings.TrimSpace(s); s != "" {
			out = append(out, s)
		}
	}
	return out
}

func originAllowed(origin string, whitelist []string) bool {
	for _, o := range whitelist {
		if o == origin {
			return true
		}
	}
	return false
}

// Cors 跨域问题
func Cors() gin.HandlerFunc {
	whitelist := allowedOrigins()
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" && originAllowed(origin, whitelist) {
			// 回显具体 Origin,不能用通配 "*":它与 Allow-Credentials: true 冲突,
			// 浏览器会直接拒绝带凭据的跨域请求。
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Vary", "Origin")
			var headerKeys []string
			for k := range c.Request.Header {
				headerKeys = append(headerKeys, k)
			}
			allowHeaders := strings.Join(headerKeys, ", ")
			if allowHeaders == "" {
				allowHeaders = "Content-Type, Authorization"
			}
			c.Header("Access-Control-Allow-Headers", allowHeaders)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		// 预检请求直接返回并中止,不再落到后面的业务 handler。
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// isTraceSkipPath 不需要链路追踪的路径(健康检查/监控/静态资源)
func isTraceSkipPath(path string) bool {
	return path == "/health" || path == "/metrics" || strings.HasPrefix(path, "/file/")
}

// UseJager 链路追踪
func UseJager(_tracer opentracing.Tracer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if isTraceSkipPath(path) { // 健康检查等请求不打点, 避免产生无用 span
			ctx.Next()
			return
		}

		// 优先接上上游传来的 span(网关注入的 trace 头), 取不到则作为 root span
		var parent opentracing.SpanContext
		if sc, err := _tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header)); err == nil {
			parent = sc
		}
		span := _tracer.StartSpan(path, ext.RPCServerOption(parent))
		defer span.Finish() // 必须 defer: handler panic 时也要 Finish, 否则 span 永远不会上报

		ext.HTTPUrl.Set(span, ctx.Request.URL.String())
		ext.HTTPMethod.Set(span, ctx.Request.Method)

		// 以 Request.Context() 为父, 并回写请求上下文: 下游取 Request.Context() 才能拿到 span
		c := opentracing.ContextWithSpan(ctx.Request.Context(), span)
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Set(api.JagerCtx, c)

		ctx.Next()

		span.SetTag("request", ctx.GetString("jager_req"))
		span.SetTag("response", ctx.GetString("jager_resp"))
		code := ctx.Writer.Status()
		ext.HTTPStatusCode.Set(span, uint16(code))
		ext.Error.Set(span, code >= http.StatusInternalServerError)
		// tracer.CloseTracer()
	}
}

// JagerGinBeforeAfter 创建一个默认 BeforeAfter Middleware
type JagerGinBeforeAfter struct {
}

type timeTrace struct{}

// GinBefore call之前调用
func (d *JagerGinBeforeAfter) GinBefore(req *ginrpc.GinBeforeAfterInfo) bool {
	// ginrpc 默认给的是 context.Background(), 换成带 span 的上下文, 前后钩子拿到的 ctx 才能接进同一条链路
	req.Context = context.WithValue(api.TraceCtx(req.C), timeTrace{}, time.Now())
	req.C.Set("jager_req", tools.JSONDecode(req.Req))
	return true
}

// GinAfter call之后调用
func (d *JagerGinBeforeAfter) GinAfter(req *ginrpc.GinBeforeAfterInfo) bool {
	// 历史实现直接 .(time.Time) 断言:GinAfter 没有配对的 GinBefore
	// (或 ginrpc 换掉了上下文)时直接 panic,一次中间件失误会把整个进程带走。
	var cost time.Duration
	if begin, ok := req.Context.Value(timeTrace{}).(time.Time); ok {
		cost = time.Since(begin)
	}
	mylog.Info(fmt.Sprintf("[middleware] call[%v] [%v]", req.FuncName, cost))

	msg := message.GetSuccessMsg()
	if req.Error != nil {
		msg = message.GetErrorMsg(message.InValidOp)
		// status.Convert 永远不会返回 nil(非 grpc 错误会被包成 Unknown),
		// 所以原来的 else 分支是死代码,且会把业务 error 的 code 覆盖成 2。
		// 只有实现了 GRPCStatus() 的错误才提取 code/message。
		if st, ok := status.FromError(req.Error); ok {
			msg.Code = int(st.Code())
			msg.Error = st.Message()
		} else {
			msg.Error = req.Error.Error()
		}
	} else {
		msg.Data = req.Resp
	}

	req.Resp = msg // 设置resp 结果
	req.C.Set("jager_resp", tools.JSONDecode(msg))

	return true
}
