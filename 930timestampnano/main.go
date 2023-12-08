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
