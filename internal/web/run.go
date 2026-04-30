package web

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

var Mux = chi.NewRouter()

func Run() {
	// 全局中间件设置
	Mux.Use(middleware.Recoverer) // 放在最外层捕获所有内部 panic
	Mux.Use(middleware.Logger)

	// 创建错误探测信封
	errChan := make(chan error, 1)

	// 协程启动
	go func() {
		if err := http.ListenAndServe(":9081", Mux); err != nil {
			errChan <- err
		}
	}()

	// 【核心处理】检查是否报错
	select {
	case err := <-errChan:
		// 如果通道里有错，说明端口启动失败（如 Address already in use）
		log.Fatalf("❌ [Web] 致命错误：端口可能被占用或权限不足 | %v", err)
	case <-time.After(100 * time.Millisecond):
		// 100ms 过去了没报错，说明端口占领成功
		fmt.Println("✅ [Web] 9081端口占领成功，底座已就绪")
		fmt.Println("🌐 [Web] 访问 http://localhost:9081")
	}

}
