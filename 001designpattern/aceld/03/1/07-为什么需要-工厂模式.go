package main

import "fmt"

// 水果类
type Fruit struct {
	Nanme string
	// 一些属性
}

func (f *Fruit) Show(name string) {
	if name == "apple" {
		fmt.Println("Show apple")
	} else if name == "banana" {
		fmt.Println("Show banana")
	}
	// 增加功能，太复杂，耦合
}

// 创建一个 Fruit 对象
func NewFruit(name string) *Fruit {
	fruit := new(Fruit)

	if name == "apple" {

	} else if name == "banana" {

	}
	// 增加功能，太复杂，耦合
	return fruit

}

// 业务层
func main() {
	apple := NewFruit("apple")
	apple.Show("apple")
}
