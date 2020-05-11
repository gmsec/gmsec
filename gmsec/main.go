package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/xxjwxc/public/dev"

	"gmsec/internal/config"
	_ "gmsec/internal/routers" // debug模式需要添加[mod]/routers 注册注解路由

	"github.com/gin-gonic/gin"
	"github.com/xxjwxc/ginrpc"
	"github.com/xxjwxc/ginrpc/api"
	"github.com/xxjwxc/public/myast"
	"github.com/xxjwxc/public/mydoc/myswagger"
	"github.com/xxjwxc/public/server"
)

// ReqTest demo struct
type ReqTest struct {
	UserName int32    `json:"user_name" binding:"required" default:"aaaaaa"` // 带校验方式
	Password []string `json:"password"`
	// 测试1111
	Hell Hello `json:"hellxxjwo"` // 测试
}

// Hello ...
type Hello struct {
	Index int
}

// Block 带注解路由(参考beego形式)
// @Router /block [post,get]
func (s *Hello) Block(c *api.Context, req *ReqTest) (*ReqTest, error) {
	fmt.Println(req)
	fmt.Println(s.Index)
	c.JSON(http.StatusOK, "ok")
	return nil, nil
}

//TestFun6 带自定义context跟已解析的req参数回调方式,err,resp 返回模式
func TestFun6(c *gin.Context, req ReqTest) (*ReqTest, error) {
	fmt.Println(req)
	//c.JSON(http.StatusOK, req)
	return &req, nil
}

// CallBack service call backe
func CallBack() {
	// swagger
	myswagger.SetHost("https://localhost:8080")
	myswagger.SetBasePath("gmsec")
	myswagger.SetSchemes(true, false)
	// -----end --

	_, modFile, isFind := myast.GetModuleInfo(0)
	outdir := ""
	if isFind {
		outdir = modFile + "/internal/routers"
	}

	base := ginrpc.New(ginrpc.WithCtx(func(c *gin.Context) interface{} {
		return api.NewCtx(c)
	}), ginrpc.WithDebug(dev.IsDev()), ginrpc.WithOutPath(outdir), ginrpc.WithGroup("xxjwxc"))
	router := gin.Default()

	h := new(Hello)
	h.Index = 123
	base.Register(router, h)                                                     // 对象注册
	router.POST("/test6", base.HandlerFunc(TestFun6))                            // 函数注册
	base.RegisterHandlerFunc(router, []string{"post", "get"}, "/test", TestFun6) // 多种请求方式注册

	router.Run(":8080")
}

func main() {
	if config.GetIsDev() || len(os.Args) == 0 {
		CallBack()
	} else {
		server.On(config.GetServiceConfig()).Start(CallBack)
	}
}
