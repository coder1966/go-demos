package main

/*

 */

import (
	"context"
	"fmt"
	"time"
)

func sub(x int) int {
	for {
		time.Sleep(time.Second)
		fmt.Print("!", x)
	}
	return 123
}
func main() {
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		subMain(ctx, i)
	}
	time.Sleep(time.Second * 20)
	fmt.Println("\n main 结束")
}

func subMain(ctx context.Context, x int) {
	go func() {
		fmt.Println("获得数据", sub(x))
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\n <-ctx.Done()", x)
			fmt.Println("subMain结束", x)
			return
		default:
			time.Sleep(time.Second)
		}
	}

}
