//go:build !dev

package api

import (
	"context"
	"testing"
)

// TestGetUserFromTokenEmpty 默认(非 dev)构建下空 token 必须一律拒绝,不能绕过认证。
// 只在 !dev 构建里断言:dev 构建按设计允许 admin 放行。
func TestGetUserFromTokenEmpty(t *testing.T) {
	if u, ok := devBypassUser(); ok || u != nil {
		t.Fatal("非 dev 构建下 devBypassUser 必须返回 (nil, false)")
	}
	if u, ok := GetUserFromToken(context.Background(), ""); ok || u != nil {
		t.Fatal("非 dev 构建下空 token 不得放行")
	}
}
