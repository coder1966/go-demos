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
	var interval time.Duration = time.Second * 20
	ts := time.Now().UnixNano()

	_, _ = interval, ts

	// sl:=(ts+int64(interval)-1)/int64(interval)
	sl := int64(interval) - ts%int64(interval)

	fmt.Println("sleep :==", ts, interval, sl)

	time.Sleep(time.Nanosecond * time.Duration(sl))

	fmt.Println("want got :==", ts+sl, time.Now().UnixNano())

	return

}
