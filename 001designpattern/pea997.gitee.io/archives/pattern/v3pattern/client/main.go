package main

import (
	"fmt"
	"godemos/001designpattern/pea997.gitee.io/archives/pattern/v3pattern/httplib"
	"os"
	"time"
)

func main() {

	opts := make([]httplib.Option, 0)

	// 下面的参数选项，可以任意组合
	opts = append(opts, httplib.WithPort(8080))
	opts = append(opts, httplib.WithTimeout(3*time.Second))

	// 重点在 opts...
	server, err := httplib.NewServer("localhost", opts...)
	if err != nil {
		fmt.Println(" error: ", err)
		panic("")
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(" error: ", err)
		os.Exit(1)
	}
}
