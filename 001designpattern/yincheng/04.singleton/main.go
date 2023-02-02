package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type Single struct {
	data int
}

var singleton *Single

// 获得 单例 的接口
func GetInterface() *Single {
	once.Do(func() { singleton = &Single{100} })
	// singleton = &Single{100} // 这样就不是单例
	return singleton
}

func main() {
	i1 := GetInterface()
	i2 := GetInterface()

	fmt.Println(i1 == i2)

}
