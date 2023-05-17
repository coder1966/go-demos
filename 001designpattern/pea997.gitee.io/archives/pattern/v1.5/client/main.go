package main

import (
	"fmt"
	"godemos/001designpattern/pea997.gitee.io/archives/pattern/v1.5/httplib"
	"os"
)

func main() {
	port := 0
	server, err := httplib.NewServer("localhost", httplib.Config{
		Port: &port,
	})
	if err != nil {
		fmt.Println(" error: ", err)
	}
	fmt.Println("server =: ", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(" error: ", err)
		os.Exit(1)
	}
}
