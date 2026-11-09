// Package api The next version of the underlying category will support automatic parsing of a single struct.
package api

import (
	"context"

	"github.com/gin-gonic/gin"
	ginapi "github.com/gmsec/goplugins/api"
	"github.com/xxjwxc/public/mylog"
)

const (
	//UserToken 用户token
	UserToken = "user_token" // 聚合了用户端token user_token
	// SessionToken session token
	SessionToken = "session_token"
	// JagerCtx 链路追踪上下文在 gin 中的 key(由 routers.UseJager 中间件写入)
	JagerCtx = "jager_ctx"
)

// TraceCtx 获取携带链路追踪 span 的上下文(无追踪时退化为请求上下文)
// 需要往下游服务继续透传链路时, 用它替换裸 context.Background()
func TraceCtx(gc *gin.Context) context.Context {
	if gc == nil {
		return context.Background()
	}
	if v, ok := gc.Get(JagerCtx); ok {
		if c, ok := v.(context.Context); ok && c != nil {
			return c
		}
	}
	if gc.Request != nil {
		return gc.Request.Context()
	}
	return context.Background()
}

// Context Wrapping gin context to custom context
type Context struct { // 包装gin的上下文到自定义context
	ginapi.Context
}

// TraceCtx 获取携带链路追踪 span 的上下文
func (c *Context) TraceCtx() context.Context {
	return TraceCtx(c.GetGinCtx())
}

// Value ginapi.Context 的基底是 context.TODO(), 这里对未命中的 key(如 span)回退到带链路追踪的上下文
func (c *Context) Value(key interface{}) interface{} {
	if v := c.Context.Value(key); v != nil {
		return v
	}
	if gc := c.GetGinCtx(); gc != nil && gc.Request != nil {
		return TraceCtx(gc).Value(key)
	}
	return nil
}

// GetUserInfo 获取用户信息
func (c *Context) GetUserInfo() (u *UserInfo, b bool) {
	accessToken, err := c.GetGinCtx().Cookie(UserToken)
	if err == nil {
		return GetUserFromToken(c.TraceCtx(), accessToken) // 带上链路追踪
	}

	accessToken, err = c.GetGinCtx().Cookie(SessionToken)
	if err == nil {
		return GetUserFromToken(c.TraceCtx(), accessToken)
	}

	mylog.ErrorString(err.Error()) // output error string
	return nil, false
}

// GetSessionToken find SessionToken by cookie
func (c *Context) GetSessionToken() string {
	sessionToken, _ := c.GetGinCtx().Cookie(SessionToken)

	return sessionToken
}

// GetClientIP get request ip
func (c *Context) GetClientIP() string {
	return c.GetGinCtx().ClientIP()
}

// NewAPIFunc default of custom handlefunc
func NewAPIFunc(c *gin.Context) interface{} {
	return &Context{*ginapi.NewCtx(c)}
}

// WidthContext 通过ctx获取
func WidthContext(ctx context.Context) *Context {
	c := &Context{}
	c.Context.Context = ctx
	c.WriteHeadToCtx(c.GetGinCtx()) // 默认推送header 到ctx里面
	return c
}
