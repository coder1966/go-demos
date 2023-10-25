package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/datakit", healthz)
	err := http.ListenAndServe(":8101", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(res http.ResponseWriter, req *http.Request) {
	//请求IP
	reqAddr := req.RemoteAddr
	reqIp := strings.Split(reqAddr, ":")[0]
	fmt.Println("请求IP:", reqIp)

	//返回Header
	resHeader := res.Header()
	for k, v := range req.Header {
		//设置返回Header内容
		resHeader.Set(k, getStr(v))
	}

	//设置环境变量
	version := "VERSION"
	os.Setenv(version, "v1.0")
	//设置环境变量VERSION至返回Header
	resHeader.Set(version, os.Getenv(version))

	//设置返回body
	fmt.Println("RETURN 200", time.Now())
	io.WriteString(res, "200")
}

func getStr(strSlice []string) string {
	str := ""
	for _, s := range strSlice {
		str = str + "," + s
	}
	return str
}
