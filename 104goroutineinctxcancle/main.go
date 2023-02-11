package main

/*
cancel() 可以传递下去，在子函数执行，全局退出。
*/

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sub(ctx, cancel)
	fmt.Println("sub() 回来了")
	time.Sleep(time.Millisecond * 2000)
	// cancel()
	time.Sleep(time.Millisecond * 200000)
}

func sub(ctx context.Context, cancel context.CancelFunc) {
	go subA(ctx, cancel)
	go subB(ctx, cancel)
	go subC(ctx, cancel)
	go subD(ctx, cancel)
}

func subA(ctx context.Context, cancel context.CancelFunc) {
	for i := 0; i < 10000; i++ {
		select {
		case <-ctx.Done():
			// 返回
			fmt.Println("subA <-ctx.Done()")
			return
		case <-time.After(time.Millisecond * 500):
			// 每 检查一次
			fmt.Print("A", i)
			continue
		}
	}
}

func subB(ctx context.Context, cancel context.CancelFunc) {
	for i := 0; i < 10000; i++ {
		select {
		case <-ctx.Done():
			// 返回
			fmt.Println("subB <-ctx.Done()")
			return
		case <-time.After(time.Millisecond * 500):
			// 每 检查一次
			fmt.Print("B", i)
			continue
		}
	}
}

func subC(ctx context.Context, cancel context.CancelFunc) {
	for i := 0; i < 10000; i++ {
		select {
		case <-ctx.Done():
			// 返回
			fmt.Println("subC <-ctx.Done()")
			return
		case <-time.After(time.Millisecond * 500):
			// 每 检查一次
			fmt.Print("C", i)
			continue
		}
	}
}

func subD(ctx context.Context, cancel context.CancelFunc) {
	time.Sleep(time.Second * 4)
	fmt.Print("#D#")
	cancel()
	return

}
