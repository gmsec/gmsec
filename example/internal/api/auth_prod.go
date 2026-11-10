//go:build !dev

package api

// devBypassUser 默认(非 dev)构建下的严格实现:空 token 一律不放行,
// 避免 GetUserFromToken 里的认证绕过被带进生产二进制。
func devBypassUser() (*UserInfo, bool) {
	return nil, false
}
