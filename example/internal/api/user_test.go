package api

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestStrToInt 覆盖纯函数的正常/非法输入。
func TestStrToInt(t *testing.T) {
	cases := map[string]int{
		"42":   42,
		"0":    0,
		"-7":   -7,
		"abc":  0, // 非法值降级为 0 并记日志
		"":     0,
		" 12 ": 0,
	}
	for in, want := range cases {
		if got := strToInt(in); got != want {
			t.Errorf("strToInt(%q) = %d, want %d", in, got, want)
		}
	}
}

// TestPostJSONWithTrace 用 httptest 起本地服务,不依赖真实 oauth2 环境。
func TestPostJSONWithTrace(t *testing.T) {
	payload := `{"token":"abc"}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("期望 POST,实际 %s", r.Method)
		}
		b, _ := io.ReadAll(r.Body)
		_, _ = w.Write(b) // 原样回显,便于断言请求体确实发出去了
	}))
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	body, err := postJSONWithTrace(ctx, srv.URL, payload)
	if err != nil {
		t.Fatalf("postJSONWithTrace 返回错误: %v", err)
	}
	if string(body) != payload {
		t.Fatalf("回显内容 = %q, want %q", string(body), payload)
	}
}

// TestPostJSONWithTraceTimeout 校验上下文超时能中断请求,不会无限等待。
func TestPostJSONWithTraceTimeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(300 * time.Millisecond)
		_, _ = w.Write([]byte("late"))
	}))
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, err := postJSONWithTrace(ctx, srv.URL, `{}`)
	if err == nil {
		t.Fatal("下游超时时期望返回错误")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") &&
		!strings.Contains(err.Error(), "Client.Timeout") {
		t.Fatalf("期望超时错误,实际: %v", err)
	}
}
