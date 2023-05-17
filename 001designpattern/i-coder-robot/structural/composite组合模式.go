package main

import "fmt"

// 部件
type Component interface {
	Traverse() // 穿过
}

type Leaf struct {
	value int
}

func NewLeaf(v int) *Leaf {
	return &Leaf{v}
}
func (l *Leaf) Traverse() {
	fmt.Println("l.value: ", l.value)
}

// 组合
type Composite struct {
	children []Component
}

func NewComposite() *Composite {
	return &Composite{make([]Component, 0)}
}

// add 加的是一个 部件 接口，实际上是一个leaf对象
func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

// 调用群体方法，每一个个体都会
func (c *Composite) Traverse() {
	for idx, _ := range c.children {
		c.children[idx].Traverse()
	}
}

func main() {
	containers := make([]Composite, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			containers[i].Add(NewLeaf(i*3 + j))
		}
	}

	for i := 0; i < 4; i++ {
		containers[0].Add(&containers[i])
	}

	for i := 0; i < 4; i++ {
		containers[i].Traverse()
		fmt.Println("finish")
	}
}
