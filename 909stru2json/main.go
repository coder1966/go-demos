package main

import (
	"fmt"
	"net/url"
)

func main() {
	urls := []string{
		"127.0.0.1",
		"http://127.0.0.1",
		"https://127.0.0.1",
		"127.0.0.1:22",
		"http://127.0.0.1:22",
		"https://127.0.0.1:22",
		"baidu.cc:22",
		"baidu.cc:22/wo/ni",
		"http://www.baidu.cc:22/wo/ni",
		"http://www.baidu.cc/wo/ni",
	}

	test01(urls)
	fmt.Println("========")

	test02(urls)
	fmt.Println("========")
}

func test01(urls []string) {
	for _, u := range urls {
		uu, err := url.Parse(u)
		fmt.Println(uu, err)
	}
}

func test02(urls []string) {
	for _, u := range urls {
		uu, err := url.ParseRequestURI(u)
		fmt.Println(uu, err)
	}
}
