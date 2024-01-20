package syncpool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// $ GODEBUG=gctrace=1 go test -benchmem -bench Benchmark_mainOld -cpu=1
// gc 428 @66.864s 9%: 0.013+50+0.002 ms clock, 0.013+3.0/12/0+0.002 ms cpu, 440->499->264 MB, 517 MB goal, 0 MB stacks, 0 MB globals, 1 P
// Benchmark_mainOld              1        67022649312 ns/op       102410455416 B/op       100003433 allocs/op
// ok      github.com/coder1966/go-demonew/a0700unittest/syncpool  67.048s
func Benchmark_mainOld(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mainOld()
	}
}

// $ GODEBUG=gctrace=1 go test -benchmem -bench Benchmark_mainPool -cpu=1
// gc 20 @47.276s 0%: 0.015+40+0.003 ms clock, 0.015+2.4/10/0+0.003 ms cpu, 409->410->206 MB, 413 MB goal, 0 MB stacks, 0 MB globals, 1 P
// Benchmark_mainPool             1        50577747975 ns/op       2615233704 B/op 100203377 allocs/op
// ok      github.com/coder1966/go-demonew/a0700unittest/syncpool  50.607s
func Benchmark_mainPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mainPool()
	}
}

// =====================================

const (
	windowSize = 200000
	msgCount   = 100000000
)

type (
	message []byte
	buffer  map[int]message
)

var worst time.Duration

func mkMessage(n int) message {
	m := make(message, 1024)
	for i := range m {
		m[i] = byte(n)
	}
	return m
}

func pushMsg(b *buffer, highID int) {
	start := time.Now()
	m := mkMessage(highID)
	(*b)[highID%windowSize] = m
	elapsed := time.Since(start)
	if elapsed > worst {
		worst = elapsed
	}
}

func mainOld() {
	b := make(buffer, windowSize)
	for i := 0; i < msgCount; i++ {
		pushMsg(&b, i)
	}
	fmt.Println("Worst push time: ", worst)
}

// =========================================

// const (
// 	windowSize = 200000
// 	msgCount   = 100000000
// )

// type (
// 	message []byte
// 	buffer  map[int]message
// )

// var worst time.Duration

// pool for statistics model
var statModelPool = sync.Pool{
	New: func() interface{} {
		return make(message, 1024)
	},
}

func mkMessagePool(n int) message {
	m := statModelPool.Get().(message)
	for i := range m {
		m[i] = byte(n)
	}
	return m
}

func pushMsgPool(b *buffer, highID int) {
	start := time.Now()
	m := mkMessagePool(highID)
	if highID > windowSize {
		statModelPool.Put((*b)[highID%windowSize])
	}

	(*b)[highID%windowSize] = m
	elapsed := time.Since(start)
	if elapsed > worst {
		worst = elapsed
	}
}

func mainPool() {
	b := make(buffer, windowSize)
	for i := 0; i < msgCount; i++ {
		pushMsgPool(&b, i)
	}
	fmt.Println("Worst push time: ", worst)
}
