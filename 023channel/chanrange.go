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
*/

import (
	"fmt"
	"time"
)

func chanRange() {
	fmt.Println("### 开始func chanRange()")
	cRangeCh := make(chan int, 10)
	go func() {
		for i := 1; i < 5; i++ {
			cRangeCh <- i
		}
		close(cRangeCh)
	}()
	for data := range cRangeCh {
		//  for data := range cRangeCh 如果  channel 关闭 就读完缓存，退出。 如果  channel 不关闭 就阻塞
		fmt.Println("data:=range :", data)
		time.Sleep(time.Second)
	}

	fmt.Println("main() 结束")

}
