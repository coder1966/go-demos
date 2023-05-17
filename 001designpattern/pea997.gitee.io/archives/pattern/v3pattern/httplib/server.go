package httplib

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var defaultPort = 8000

type Options struct {
	port    *int
	timeout time.Duration
}

type Option func(option *Options) error

// WithPort 准备被传入的参数，实际上是一个函数
func WithPort(port int) Option {
	// 返回一个匿名函数，是一个闭包。这个匿名函数里面，执行我要的逻辑。
	return func(options *Options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}

}

func WithTimeout(timeout time.Duration) Option {
	// 返回一个匿名函数，是一个闭包。这个匿名函数里面，执行我要的逻辑。
	return func(options *Options) error {
		options.timeout = timeout
		return nil
	}
}

//  NewServer 每个 Option 都是一个 func
func NewServer(addr string, opts ...Option) (*http.Server, error) {
	var options Options // 真正的配置选项

	for _, opt := range opts {
		// 执行每一个 opts 对应的方法，修正配置选项 options 。options传参进
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}

	var port int
	if options.port == nil {
		port = defaultPort
	}
	if options.port != nil {
		if *options.port == 0 {
			port = randomPort()
		} else {
			port = *options.port
		}
	}

	return &http.Server{Addr: fmt.Sprintf("%s:%d", addr, port)}, nil

}

func randomPort() int {
	return 1234
}
