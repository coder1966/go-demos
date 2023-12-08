package main

import (
	"io"
	"log"
	"net/http"
)

// 假设需要测试某个 API 接口的 handler 能够正常工作，例如 helloHandler

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		log.Printf("reponse err ")
	}
}
