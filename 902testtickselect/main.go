package main

import (
	"fmt"
	"time"
)

var stopAllCh chan struct{} = make(chan struct{}) // If put any in this channel all goroutine will cancel.

func main() {

	fmt.Println("\nmain 开始")
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	// 这个是错误的示范
	select {
	case <-tick.C:
		fmt.Print("打点")
		<-stopAllCh // Blocking, wait anyone send stopAll.
		fmt.Println("\n获得关闭<-stopAllCh")
	default:
		fmt.Println("\n走了default")
	}

	for i := 0; i < 10; i++ {
		fmt.Print("循环", i)
		time.Sleep(time.Second * 2)
	}

	fmt.Println("\nmain 结束")
}
