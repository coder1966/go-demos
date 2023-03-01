package main

/*
	用 tick 比 sleep 的优点可以立刻得到反馈
*/

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	go delayCloseCh(stopCh)
	sub(stopCh)
	fmt.Println("main 结束")
}
func delayCloseCh(stopCh chan struct{}) {
	fmt.Print("10秒后关闭通道")
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Print(i)
	}
	fmt.Println("关闭通道")
	close(stopCh)
}
func sub(stopCh chan struct{}) {
	defer fmt.Println("退出了sub()函数")
	tick := time.NewTicker(time.Second * 8)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
		case <-stopCh:
			return
		}
		fmt.Println("每停顿8秒执行一次的工作")
	}
}
