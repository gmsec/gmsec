package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"example/internal/config"

	"github.com/opentracing/opentracing-go"
	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/mycache"
	"github.com/xxjwxc/public/mylog"
)

// checkTokenTimeout 校验 token 的下游调用超时,避免没有超时的请求把服务拖死。
const checkTokenTimeout = 5 * time.Second

// checkTokenClient 复用带超时的 Client,替代无超时的 http.DefaultClient。
var checkTokenClient = &http.Client{Timeout: checkTokenTimeout}

// UserInfo 用户信息
type UserInfo struct {
	AccessToken string // token
	UserName    string // 用户名
	ExpireTime  int    // 超时时间
}

// MapMessageBody req message body
type MapMessageBody struct {
	message.MessageBody
	Data map[string]string `json:"data,omitempty"`
}

// GetUserFromToken 通过token获取用户信息
// ctx 需携带链路追踪 span(可用 TraceCtx 获取), 以便下游调用接入同一条链路
func GetUserFromToken(ctx context.Context, token string) (userInfo *UserInfo, b bool) {
	if ctx == nil {
		// 裸 context.Background() 没有超时,下游 hang 住会一直等待,这里补上超时。
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), checkTokenTimeout)
		defer cancel()
	}
	if len(token) == 0 {
		// 认证绕过只在 -tags dev 构建下存在(见 auth_dev.go / auth_prod.go),
		// 默认(生产)构建恒返回 (nil, false)。
		return devBypassUser()
	}

	// 先从缓存中获取
	if tmp, ok := GetCacheBody(token); ok {
		return &UserInfo{
			AccessToken: token,
			UserName:    tmp.UserName,
			ExpireTime:  tmp.ExpireTime}, true
	}

	parm := make(map[string]string)
	parm["token"] = token
	bod, err := json.Marshal(parm)
	if err != nil {
		mylog.Error(err)
		return
	}
	rBody, err := postJSONWithTrace(ctx, config.GetCheckTokenURL(), string(bod))
	if err != nil {
		mylog.Error(err)
		return
	}
	if len(rBody) > 0 {
		var msg MapMessageBody
		if err := json.Unmarshal(rBody, &msg); err != nil {
			mylog.Error(err)
			return
		}
		if msg.State {
			SaveCacheBody(token, msg.Data["username"], msg.Data["expire_time"]) // 保存缓存
			return &UserInfo{
				AccessToken: token,
				UserName:    msg.Data["username"],
				ExpireTime:  strToInt(msg.Data["expire_time"])}, true // 返回结果
		}
	}

	return
}

// postJSONWithTrace 带链路追踪的 POST 请求(把当前 span 注入到请求头, 无 trace 上下文时等价于普通请求)
func postJSONWithTrace(ctx context.Context, url, jsonStr string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	if span := opentracing.SpanFromContext(ctx); span != nil { // 下游服务据此把 span 挂到同一条链路上
		if err = span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header)); err != nil {
			mylog.Error(err)
		}
	}

	resp, err := checkTokenClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// SaveCacheBody 保存缓存
func SaveCacheBody(accessToken, username, expireTime string) {
	tmp := UserInfo{
		AccessToken: accessToken,
		UserName:    username,
		ExpireTime:  strToInt(expireTime),
	}
	// 保存缓存
	cache := mycache.NewCache("oauth2")
	cache.Add(accessToken, &tmp, time.Duration(tmp.ExpireTime)*time.Second)
	//------------------end
}

// GetCacheBody 获取缓存
func GetCacheBody(token string) (value *UserInfo, b bool) {
	value = new(UserInfo)
	cache := mycache.NewCache("oauth2")
	err := cache.Value(token, &value)
	if err != nil {
		return nil, false
	}
	return value, true
}

func strToInt(src string) int {
	i, err := strconv.Atoi(src)
	if err != nil {
		// 原来直接丢弃 err,非法值会静默变 0;至少记日志,便于排查下游脏数据。
		mylog.Errorf("strToInt: 解析 %q 失败: %v", src, err)
		return 0
	}
	return i
}
