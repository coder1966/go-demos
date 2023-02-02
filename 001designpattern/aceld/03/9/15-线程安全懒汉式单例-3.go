package main

import (
	"sync"
)

var onec sync.Once

type singelton struct{}

var instance *singelton

// func GetInstance() *singelton {
// 	// 只有首次调用，才会。多线程，可能同时来
// 	if instance == nil {
// 		instance = new(singelton)
// 		return instance
// 	}
// 	return instance
// }

func GetInstance() *singelton {

	onec.Do(func() {
		instance = new(singelton)
	})

	return instance
}

func main() {

}
