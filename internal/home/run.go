package home

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"sbcc/internal/web"
)

func Run() {
	// 模拟后加载逻辑

	r := chi.NewRouter()
	// ... 这里可以无脑复制官方文档的代码 ...
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Home!"))
	})

	// 最后一炮打到底座，搞定！
	web.Mux.Mount("/", r)
	fmt.Println("🏠 [Home] 主页模块 加载完成！")

}
