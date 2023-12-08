package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("===== start main =====")
	// 路由与视图函数绑定
	http.HandleFunc("/health-check", HealthCheckHandler)

	// 启动服务,监听地址
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("===== end main =====")
}
