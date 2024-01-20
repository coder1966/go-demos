package main

import (
	"fmt"
	"time"
)

// func visit(path string, di fs.DirEntry, err error) error {
// 	fmt.Printf("Visited: %s\n", path)
// 	return nil
// }

func main() {
	getData()
	var interval time.Duration = time.Millisecond * 100
	var interval2 time.Duration = time.Millisecond * 1000

	tick := time.NewTicker(interval)
	defer tick.Stop()
	defer tick.Stop()

	for i := 0; i < 10; i++ {

		fmt.Println("now = ", time.Now().UnixMicro())
		if 2 < i && i < 6 {
			tick = time.NewTicker(interval2)
		}
		if i >= 6 {
			tick.Reset(interval)
		}
		<-tick.C
	}
}

/*
延迟超过200毫秒的，就认定为新一族的数据
*/

func getData() {
	ch := make(chan int, 100)
	go send(ch)

	tick := time.NewTicker(time.Millisecond * 200)
	defer tick.Stop()

	d := []int{}
	for {
		select {
		case v, ok := <-ch:
			fmt.Println("=== ", v)
			if !ok {
				fmt.Println("get === ", d)
				return
			}
			d = append(d, v)
			tick.Reset(time.Millisecond * 200)
		case <-tick.C:
			fmt.Println("get === ", d)
			d = []int{}
		}
	}
}

func send(ch chan int) {
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6}
	for _, v := range d {
		ch <- v
		if v == 0 {
			time.Sleep(time.Millisecond * 210)
		} else {
			time.Sleep(time.Millisecond * 190)
		}
		fmt.Println("send --- ", v)
	}
	close(ch)
}
