/*
空调 电视 扫地机器人 遥控器 分散
遥控器 外观不统一

中间  防火城 中间层 接口


package main
// ###### 抽象层
// ###### 基础模块层
// ###### 业务层
func main() {
}

*/

package main

import "fmt"

type SubSysA struct{}

func (sa *SubSysA) MethodA() {
	fmt.Println("子系统A 的 方法 A")
}

type SubSysB struct{}

func (sb *SubSysB) MethodB() {
	fmt.Println("子系统B 的 方法 B")
}

type SubSysC struct{}

func (sc *SubSysC) MethodC() {
	fmt.Println("子系统C 方法 C")
}

// 外观模式，提供一个外观，吧ABC 包裹进来， 简化成一个简单的接口 组合调用
type Facade struct {
	a *SubSysA
	b *SubSysB
	c *SubSysC
}

// 增加各种组合
func (f *Facade) Method01() {
	f.a.MethodA()
	f.b.MethodB()
}

func (f *Facade) Method02() {
	f.a.MethodA()
	f.c.MethodC()
}

func main() {
	sa := new(SubSysA)
	sa.MethodA()

	sb := new(SubSysB)
	sb.MethodB()

	// 使用外观模式 包裹能力

	f := Facade{
		a: new(SubSysA),
		b: new(SubSysB),
		c: new(SubSysC),
	}

	// 调用外观包裹
	f.Method01()
	f.Method02()

}
