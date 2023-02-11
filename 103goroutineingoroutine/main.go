package main

/*

 */

import (
	"fmt"
	"time"
)

func main() {
	sub()
	fmt.Println("sub() 回来了")
	time.Sleep(time.Millisecond * 200000)
}

func sub() {
	go sub02()
}

func sub02() {
	for i := 0; i < 100; i++ {
		fmt.Print(",", i)
		time.Sleep(time.Millisecond * 200)
	}
}
