package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("结束前执行")
	fmt.Println("开始执行")
	test()
	fmt.Println("恐慌后执行的代码")
}

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover() 捕获错误 error: ", err)
		} else {
			fmt.Println("recover() 没有捕获错误 error: ", err)
		}
	}()
	n1 := 0
	n2 := 4
	_ = n2 / n1 // 产生 恐慌
}
