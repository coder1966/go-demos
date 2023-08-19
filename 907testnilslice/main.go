package main

import (
	"fmt"
	"strings"

	"github.com/dustin/go-humanize"
)

func main() {
	var nilStrs []string
	var empStrs []string
	nilStrs = nil
	empStrs = make([]string, 0)

	fmt.Println("nilStrs : ", nilStrs)
	fmt.Println("empStrs : ", empStrs)

	fmt.Println("lennilStrs : ", len(nilStrs))
	fmt.Println("lenempStrs : ", len(empStrs))

	fmt.Println("nilStrs==nil : ", nilStrs == nil)
	fmt.Println("empStrs==nil : ", empStrs == nil)

	for _, v := range nilStrs {
		_ = v
	}
}

type writeCounter struct {
	total   uint64
	current uint64
	last    float64
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.current += uint64(n)
	wc.last += float64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc *writeCounter) PrintProgress() {
	if wc.last > float64(wc.total)*0.01 || wc.current == wc.total { // update progress-bar each 1%
		fmt.Printf("\r%s", strings.Repeat(" ", 36)) //nolint:gomnd
		fmt.Printf("\rDownloading(% 7s)... %s/%s", CurDownloading, humanize.Bytes(wc.current), humanize.Bytes(wc.total))
		wc.last = 0.0
	}
}
