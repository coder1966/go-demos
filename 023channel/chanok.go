package main

/*
data, ok := <-
读未关闭，有数据，就data+true
		无数据，就阻塞
读关闭，有数据，就data+true
		无数据，就0值+false

向关闭的 channel 发数据会 panic

向 nil 的(没有 make 的) ， channel 发数据会 阻塞，不是 panic
*/

import (
	"fmt"
	"time"
)

func chanOk() {
	fmt.Println("### 开始func chanOk()")
	/* cOkCh := make(chan int)

	go func() {
		for i := 1; i < 5; i++ {
			cOkCh <- i
		}
	}()

	for {
		// 如果 ok，就是 channel 没有关闭
		if data, ok := <-cOkCh; ok {
			fmt.Println("ok=true :", data)
		} else {
			fmt.Println("ok=false :", data)
		}
	} */

	cOkCh := make(chan int, 10)
	// var cOkCh chan int // 向 nil 的(没有 make 的) ， channel 发数据会 阻塞，不是 panic
	go func() {
		for i := 1; i < 5; i++ {
			cOkCh <- i
		}
		close(cOkCh)
	}()

	for {
		// 如果 ok，就是 channel 没有关闭
		if data, ok := <-cOkCh; ok {
			fmt.Println("ok=true :", data)
		} else {
			fmt.Println("ok=false :", data)
			break
		}
		time.Sleep(time.Second)
	}

	fmt.Println("main() 结束")

}
