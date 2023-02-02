package main

import (
	"sync"
	"sync/atomic"
)

type singelton struct{}

var instance *singelton

var lock sync.Mutex

var initialized uint32

// func GetInstance() *singelton {
// 	// 只有首次调用，才会。多线程，可能同时来
// 	if instance == nil {
// 		instance = new(singelton)
// 		return instance
// 	}
// 	return instance
// }

func GetInstance() *singelton {
	// 只有首次调用，才会。
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	// 没有初始化
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = new(singelton)
		// 设置标记为
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func main() {

}
