package main

import (
	_ "gmsec/routers" // debug模式需要添加[mod]/routers 注册注解路由
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/xxjwxc/ginrpc"
	"github.com/xxjwxc/ginrpc/api"
)

// TestMain test main
func TestMain(t *testing.T) {

	// debug test
	// wp := workpool.New(1000)    // Set the maximum number of threads
	// for i := 0; i < 2000; i++ { // Open 20 requests
	// 	wp.Do(func() error {
	// 		orm := core.Dao.GetDBr()
	// 		var ut []model.UserInfoTbl
	// 		err := orm.Table("user_info_tbl").Find(&ut).Error
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 			return err
	// 		}
	// 		fmt.Println(tools.GetJSONStr(ut, true))
	// 		return nil
	// 	})
	// }

	// wp.Wait()
	// fmt.Println("down")

	// swagger
	// -----end --

	base := ginrpc.New(ginrpc.WithCtx(func(c *gin.Context) interface{} {
		return api.NewCtx(c)
	}), ginrpc.WithDebug(true), ginrpc.WithGroup("xxjwxc"))
	router := gin.Default()

	h := new(Hello)
	h.Index = 123
	base.Register(router, h)                                                     // 对象注册
	router.POST("/test6", base.HandlerFunc(TestFun6))                            // 函数注册
	base.RegisterHandlerFunc(router, []string{"post", "get"}, "/test", TestFun6) // 多种请求方式注册

	//router.Run(":8080")
}
