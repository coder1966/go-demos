package main

/*
	main 调用 A
	A 调用 B
	A 退出 ，B 并不退出
*/

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("\n main 开始")

	go subA()

	time.Sleep(time.Second * 8)
	fmt.Println("\n main 结束")
}

func subA() {
	fmt.Println("\n subA 开始")
	go subB()
	fmt.Println("\n subA 结束")

}

func subB() {
	fmt.Println("\n subB 开始")
	for i := 0; i < 12; i++ {
		fmt.Print("|B", i)
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("\n subB 结束")
}
