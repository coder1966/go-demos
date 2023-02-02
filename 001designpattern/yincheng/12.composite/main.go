package main

import "fmt"

// 设计一个代理
type Subject interface {
	Do() string // 实际业务。例如几个系统的挂接。业务：检查欠费，检查密码正确
}

type RealSubject struct{}

func (s RealSubject) Do() string {
	return "真正对象 放回的"
}

type Proxy struct {
	real  RealSubject // 代理把业务包装进来
	money int         // 包赚，额外加 money passwd 等
}

func (p Proxy) Do() string {
	if p.money > 0 {
		return p.real.Do() // 钱包够，就干活
	} else {
		return "欠费，需要充值"
	}
}

func main() {
	var sub Subject
	sub = &Proxy{money: -1}
	fmt.Println(sub.Do())

	sub = &Proxy{money: 1}
	fmt.Println(sub.Do())
}
