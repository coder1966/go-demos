package main

import "fmt"

/*
concreteBuilder 具体构建者
Builder 构建各个部分 buildPart()
director 导演
*/

type Builder interface {
	buildDisk()
	buildCPU()
	buildRom()
}

// 一个型号的电脑
type SuperComputer struct{ Name string }

func (c *SuperComputer) buildDisk() { fmt.Println("超大硬盘") }
func (c *SuperComputer) buildCPU()  { fmt.Println("超快CPU") }
func (c *SuperComputer) buildRom()  { fmt.Println("超大内存") }

// 另外一个型号的电脑
type LowComputer struct{ Name string }

func (c *LowComputer) buildDisk() { fmt.Println("小硬盘") }
func (c *LowComputer) buildCPU()  { fmt.Println("小CPU") }
func (c *LowComputer) buildRom()  { fmt.Println("小内存") }

// 导演 关联 创建者
type Director struct {
	builder Builder
}

//  NEW 一个 建造者，传入 建造者，返回 导演，导演包含建造者
func NewConstruct(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

// 导演，构建，一下三个出来了
func (d *Director) Construct() {
	d.builder.buildDisk()
	d.builder.buildCPU()
	d.builder.buildRom()

}

// 主程序
func main() {

	sc := SuperComputer{}
	d := NewConstruct(&sc)
	d.Construct()

	lc := LowComputer{}
	d2 := NewConstruct(&lc)
	d2.Construct()

}
