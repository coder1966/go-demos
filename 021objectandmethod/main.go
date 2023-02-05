package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Age  int
}

func (a *Animal) Eat() {
	fmt.Println("动物 吃饭")
}
func (a *Animal) Say() {
	fmt.Println("动物 吼叫")
}

// Human 继承动物的类和方法
type Human struct {
	Animal
}

// Say 重新定义父类的 Say() 方法，不重新定义 Eat()
func (h *Human) Say() {
	fmt.Println("人类 说话")
}

func main() {
	a := Animal{}
	h := Human{}

	a.Eat()
	a.Say()
	h.Eat()
	h.Say()

}
