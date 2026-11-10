//go:build dev

package api

// devBypassUser 仅在使用 -tags dev 构建时编译:空 token 放行 admin,方便本地联调。
// 严禁在生产/默认构建里启用,见 auth_prod.go。
func devBypassUser() (*UserInfo, bool) {
	return &UserInfo{
		UserName: "admin",
	}, true
}
