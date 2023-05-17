package main

import (
	"fmt"
	"godemos/001designpattern/pea997.gitee.io/archives/pattern/v2builder/httplib"
	"os"
	"time"
)

func main() {
	builder := httplib.ConfigBuilder{}
	// cfg, err := builder.Port(0).Build()
	cfg, err := builder.Port(0).Timeout(2 * time.Second).Build()
	if err != nil {
		fmt.Println(" error: ", err)
		panic("")
	}

	server, err := httplib.NewServer("localhost", cfg)
	if err != nil {
		fmt.Println(" error: ", err)
		panic("")
	}
	fmt.Println("server =: ", server.Addr)

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(" error: ", err)
		os.Exit(1)
	}
}
