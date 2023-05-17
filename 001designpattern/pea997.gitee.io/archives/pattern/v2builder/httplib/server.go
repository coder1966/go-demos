package httplib

import (
	"fmt"
	"net/http"
	"time"
)

var defaultPort = 8000

type Config struct {
	Port    int
	Timeout time.Duration
}

// 建造者
type ConfigBuilder struct {
	port    *int
	timeout time.Duration
}

// Port 暴露 Port 方法。回 建造者，有利于链式调用
func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b
}

// Timeout 增加暴露 Timeout 方法。回 建造者，有利于链式调用
func (b *ConfigBuilder) Timeout(timeout time.Duration) *ConfigBuilder {
	b.timeout = timeout
	return b
}

// Build 暴露 建造 方法。得到配置文件
func (b *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}

	if b.port == nil {
		cfg.Port = defaultPort
	} else {
		if *b.port == 0 {
			cfg.Port = randomPort()
		} else if *b.port < 0 {
			return Config{}, fmt.Errorf("pot cannot be negative ")
		} else {
			cfg.Port = *b.port
		}
	}

	return cfg, nil
}

func NewServer(addr string, cfg Config) (*http.Server, error) {

	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", addr, cfg.Port),
	}
	return server, nil
}

func randomPort() int {
	return 1234
}
