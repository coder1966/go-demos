package main

/*
data, ok := <-
读未关闭，有数据，就data+true
		无数据，就阻塞
读关闭，有数据，就data+true
		无数据，就0值+false

向关闭的 channel 发数据会 panic

向 nil 的(没有 make 的) ， channel 发数据会 阻塞，不是 panic

for data := range cRangeCh 如果  channel 关闭 就读完缓存，退出。 如果  channel 不关闭 就阻塞
*/
import "fmt"

var intChan chan int

func main() {
	main01()
	chanOk()
	chanRange()
	chanSelect()
}

func main01() {

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
