package main

import (
	"fmt"
	"godemos/001designpattern/pea997.gitee.io/archives/pattern/v1/httplib"
	"os"
)

func main() {
	server, err := httplib.NewServer("localhost", 8000)
	if err != nil {
		fmt.Println(" error: ", err)
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(" error: ", err)
		os.Exit(1)
	}
}
