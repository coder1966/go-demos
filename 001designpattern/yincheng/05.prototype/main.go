package main

import "fmt"

// 浅拷贝，深拷贝

// 原型对象需要实现的接口 ###### 返回的是自己
// ###### 特色就是把对象复制自己
type Cloneable interface {
	Clone() Cloneable
}

// 原型对象的类
type PrototypeManger struct {
	prototypes map[string]Cloneable
}

// 构造函数
func NewPrototypeManger() *PrototypeManger {
	return &PrototypeManger{make(map[string]Cloneable)}
}
func (p *PrototypeManger) Get(name string) Cloneable {
	return p.prototypes[name]
}
func (p *PrototypeManger) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

//  ==== 代表第一个类型  目的，深复制|浅复制
type Type01 struct {
	name string
}

func (t *Type01) Clone() Cloneable {
	// tc := *t // 拷贝内容，深拷贝。开辟内存新建变量，存储指针执行的内容
	// return &tc // 返回的是地址
	return t // 浅拷贝
}

//  ==== 代表第一个类型  目的，不同类型拷贝
type Type02 struct {
	name string
}

func (t *Type02) Clone() Cloneable {
	tc := *t   // 拷贝内容，深拷贝。开辟内存新建变量，存储指针执行的内容
	return &tc // 返回的是地址
}

func main() {
	// 初始化
	mgr := NewPrototypeManger()
	// 创建对象
	t1 := &Type01{name: "type1"}

	mgr.Set("t01", t1)
	t11 := mgr.Get("t01")
	t22 := t11.Clone()
	if t11 == t22 {
		fmt.Println("浅拷贝")
	} else {
		fmt.Println("深拷贝")
	}

}
