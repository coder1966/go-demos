package main

/*
data, ok := <-
读未关闭，有数据，就data+true
		无数据，就阻塞
读关闭，有数据，就data+true
		无数据，就0值+false

向关闭的 channel 发数据会 panic

向 nil 的(没有 make 的) ， channel 发数据会 阻塞，不是 panic

for data := range cRangeCh 如果  channel 关闭 就读完缓存，退出。 如果  channel 不关闭 就阻塞

select 监控多个 chan ，哪个不阻塞就执行哪个，都阻塞执行 default
*/

import (
	"fmt"
	"time"
)

func chanSelect() {
	fmt.Println("### 开始func chanSelect()")
	c01 := make(chan int, 3)
	c02 := make(chan int, 3)
	go func() {
		for i := 1; i < 5; i++ {
			c01 <- i
			time.Sleep(time.Second * 2)
		}
		// close(c01)
	}()
LOOP:
	for {
		select {
		case data := <-c01:
			fmt.Println("case data <- c01: ", data)
		case c02 <- 555:
			fmt.Println("执行了 case c02 <- 555:")
		default:
			fmt.Println("执行了 default:  break LOOP")
			break LOOP
		}
		time.Sleep(time.Second)
	}

	fmt.Println("main() 结束")

}
