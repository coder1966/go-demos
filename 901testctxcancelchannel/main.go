package main

import (
	"context"
	"fmt"
	"time"
)

var stopAllCh chan struct{} = make(chan struct{}) // If put any in this channel all goroutine will cancel.

func main() {
	// 用 onece

	fmt.Println("\nmain 一半")
	// 不用 onece
	for i := 0; i < 10; i++ {
		// 必须每次重新创建ctx
		retryCtx, retryCancel := context.WithCancel(context.Background())
		go func() {
			<-stopAllCh   // 阻塞，等待任何人发信号
			retryCancel() // 发信号停止
		}()
		sub(retryCtx)
		<-retryCtx.Done() // 阻塞一下
		fmt.Println("\n主函数循环", i)
	}
	fmt.Println("\nmain 结束")
}

func sub(retryCtx context.Context) {
	go subA(retryCtx)
	go subB(retryCtx)
	go subC(retryCtx)
	go subD(retryCtx)
}

func subA(retryCtx context.Context) {
	defer fmt.Println("\n停止了 A")
	for i := 0; i < 10000; i++ {
		select {
		case <-retryCtx.Done():
			// 父进程发停止了
			return
		case <-time.After(time.Millisecond * 500):
			// 每 检查一次
			fmt.Print("检查 A", i)
			continue
		}
	}
}

func subB(retryCtx context.Context) {
	defer fmt.Println("\n停止了 B")
	for i := 0; i < 10000; i++ {
		select {
		case <-retryCtx.Done():
			// 父进程发停止了
			return
		case <-time.After(time.Millisecond * 500):
			// 每 检查一次
			fmt.Print("检查 B", i)
			continue
		}
	}
}

func subC(retryCtx context.Context) {
	defer fmt.Println("\n停止了 C")
	for i := 0; i < 10000; i++ {
		select {
		case <-retryCtx.Done():
			// 父进程发停止了
			return
		case <-time.After(time.Millisecond * 500):
			// 每 检查一次
			fmt.Print("检查 C", i)
			continue
		}
	}
}

func subD(retryCtx context.Context) {
	defer fmt.Println("\n停止了 D")

	for i := 0; i < 10000; i++ {

		if i > 10 {
			// 这里人为关闭
			stopAllCh <- struct{}{}
		}

		select {
		case <-retryCtx.Done():
			// 父进程发停止了
			return
		case <-time.After(time.Millisecond * 500):
			// 每 检查一次
			fmt.Print("检查 D", i)
			continue
		}

	}
}
