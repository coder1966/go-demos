package main

import (
	"fmt"
)

func main() {
	highCh := make(chan int)
	lowCh := make(chan int)

	select {
	case <-highCh:
		fmt.Println("执行了高优先级")
	case <-lowCh:

		select {
		case <-highCh:
			fmt.Println("优先执行了高优先级")
		default:
			break
		}

		fmt.Println("执行了低优先级")
	}

}
