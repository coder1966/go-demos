package main

/*
stopAllCh := make(chan struct{}) // If close this chanel all goroutine will cancel
If close this chanel all goroutine will cancel
If close this chanel all goroutine will cancel
If close this chanel all goroutine will cancel
*/

import (
	"fmt"
	"time"
)

// this is main func.
func main() {
	stopAllCh := make(chan struct{}) // If close this chanel all goroutine will cancel.
	Sub(stopAllCh)
	fmt.Println("sub() 回来了") // do it
	time.Sleep(time.Millisecond * 200000)
}

// Sub this is sub func.
func Sub(stopAllCh chan struct{}) {
	go subA(stopAllCh)
	go subB(stopAllCh)
	go subC(stopAllCh)
	go subD(stopAllCh)
}

func subA(stopAllCh chan struct{}) {
	for i := 0; i < 10000; i++ {
		select {
		case <-stopAllCh:
			// 返回
			fmt.Println("subA <-stopAllCh")
			return
		case <-time.After(time.Millisecond * 500):
			// 每 检查一次
			fmt.Print("A", i)
			continue
		}
	}
}

func subB(stopAllCh chan struct{}) {
	for i := 0; i < 10000; i++ {
		select {
		case <-stopAllCh:
			// 返回
			fmt.Println("subB <-stopAllCh")
			return
		case <-time.After(time.Millisecond * 500):
			// 每 检查一次
			fmt.Print("B", i)
			continue
		}
	}
}

func subC(stopAllCh chan struct{}) {
	for i := 0; i < 10000; i++ {
		select {
		case <-stopAllCh:
			// 返回
			fmt.Println("subC <-stopAllCh")
			return
		case <-time.After(time.Millisecond * 500):
			// 每 检查一次
			fmt.Print("C", i)
			continue
		}
	}
}

func subD(stopAllCh chan struct{}) {
	time.Sleep(time.Second * 4)
	fmt.Print("#D#")
	close(stopAllCh)
}
