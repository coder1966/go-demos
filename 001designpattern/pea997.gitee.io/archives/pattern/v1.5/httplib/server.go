package httplib

import (
	"fmt"
	"net/http"
)

type Config struct {
	Port *int
}

var defaultPort = 8000

func NewServer(addr string, cfg Config) (*http.Server, error) {
	// 默认
	if cfg.Port == nil {
		cfg.Port = &defaultPort
	}
	// 错误值
	if *cfg.Port < 0 {
		return nil, fmt.Errorf("pot cannot be negative like: %d", *cfg.Port)
	}
	server := &http.Server{}
	// 随机
	if *cfg.Port == 0 {
		server.Addr = fmt.Sprintf("%s:%d", addr, randomPort())
		return server, nil
	}
	// 通常
	server.Addr = fmt.Sprintf("%s:%d", addr, *cfg.Port)
	return server, nil
}

func randomPort() int {
	return 1234
}
