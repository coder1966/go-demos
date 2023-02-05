package main

import "fmt"

var intChan chan int

func main() {

	intChan = make(chan int, 10)
	intChan <- 1
	intChan <- 2
	intChan <- 1
	intChan <- 4
	fmt.Println("intChan 值+指针地址", <-intChan, intChan)
	fmt.Println("intChan 长度+容量", len(intChan), cap(intChan))

	// =======================
	allChan := make(chan interface{}, 10)

	allChan <- "一个"
	allChan <- 2
	allChan <- "333"

	close(allChan)

	for v := range allChan {
		fmt.Println("for range :", v)
	}

	// =======================
	allChan2 := make(chan interface{}, 10)

	allChan2 <- "一个2"
	allChan2 <- 22
	allChan2 <- "3332"

	close(allChan2)

	for {
		v, ok := <-allChan2
		if !ok {
			break
		}
		fmt.Println("for :", v)
	}

	// =======================

	intChan = make(chan int, 10)
	exitChan := make(chan struct{})
	go write(intChan)
	go read(intChan, exitChan)
	<-exitChan

	fmt.Println("携程 chan 读写 完成")

}

func write(iCh chan int) {

	for i := 0; i < 30; i++ {
		iCh <- i + 100
		fmt.Println("写入：", i+100)
	}
	close(iCh)
}
func read(iCh chan int, eCh chan struct{}) {
	for v := range iCh {
		fmt.Println("读出：", v)
	}
	eCh <- struct{}{}
	close(eCh)
}
