package main

import (
	"fmt"
	"time"
)

var ch chan int

func main() {
	ch = make(chan int, 3)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	s := make([]int, 0)
	for v := range ch {
		s = append(s, v)
		fmt.Println("获得", v)
	}
	fmt.Println("main 结束", s)

}
