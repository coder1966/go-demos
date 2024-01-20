package main

import (
	"fmt"
	"sync"
	"time"
)

/*
https://blog.51cto.com/ghostwritten/5345060
使用 Mutex 常见的错误场景有 4 类，分别是 ​​Lock/Unlock​​ 不是成对出现、Copy 已使用的 Mutex、重入和死锁

 ​​Lock/Unlock​​ 不是成对出现、
 Copy 已使用的 Mutex、
 func main() {
    var c Counter
    c.Lock()
    defer c.Unlock()
    c.Count++
    foo(c) // 复制锁
}
// 这里Counter的参数是通过复制的方式传入的
func foo(c Counter) {
    c.Lock()
    defer c.Unlock()
    fmt.Println("in foo")
}
第 12 行在调用 foo 函数的时候，调用者会复制 ​​Mutex​​​ 变量 c 作为 foo 函数的参数，不幸的是，复制之前已经使用了这个锁，这就导致，复制的 Counter 是一个带状态 ​​Counter​​。


 重入
func foo(l sync.Locker) {
    fmt.Println("in foo")
    l.Lock()
    bar(l)
    l.Unlock()
}


func bar(l sync.Locker) {
    l.Lock()
    fmt.Println("in bar")
    l.Unlock()
}


func main() {
    l := &sync.Mutex{}
    foo(l)
}
写完这个 Mutex 重入的例子后，运行一下，你会发现类似下面的错误。程序一直在请求锁，但是一直没有办法获取到锁，结果就是 Go 运行时发现死锁了，


 和死锁。
*/

func main() {
	var wg sync.WaitGroup
	//定义一个变量一个锁
	var (
		mtx1 sync.Mutex
		a    int
	)

	wg.Add(1)
	go func() {
		mtx1.Lock() //加个锁
		a = 114
		fmt.Println("[1]-->: ", a) //输出一下设定的a的值

		//睡眠5s后再输出一下a的值，因为在2s后a的值已经在主线程中被更改了
		time.Sleep(5 * time.Second)
		fmt.Println("[2]-->: ", a)

		mtx1.Unlock() //解锁
		wg.Done()
	}()

	//睡眠2s后更改一下a的值
	time.Sleep(2 * time.Second)
	a = 514

	wg.Wait()
}

/*
输出为：

[root@vm10-0-0-63 gopath]# go run a.go
[1]-->:  114
[2]-->:  514


疑问：“诶我不是加了锁了么，为什么a的值还会变？”

然后发现，噢，加了锁后不是说中间的内容不能变，锁也没那么智能会去分析你加锁后到底修改了哪些东西。

锁就只能保证，如果某个锁加锁了，那么别的地方再碰到这个锁就会阻塞，没碰到，那该干什么还干什么。
*/
