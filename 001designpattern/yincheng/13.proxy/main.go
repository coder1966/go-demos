package main

import (
	"fmt"
	"strconv"
)

// 适配的目标接口
type Target interface {
	Request(int, int) string
}
type adapter struct {
	Adaptee
}

func NewAdaptor(adaptee Adaptee) Target {
	return &adapter{Adaptee: adaptee}
}

// 接口包装，这一步实现了接口的适配
func (ad *adapter) Request(a, b int) string {
	return ad.SpecficRequest(a, b) // 故意调用另外一个
}

// 被适配
type Adaptee interface {
	SpecficRequest(int, int) string
}

// 载体,被适配的目标类 ，为了New函数能返回
type adapeeImpl struct{}

// 实际函数,翻译，桥接 作用
func (ad *adapeeImpl) SpecficRequest(a, b int) string {
	return "SpecficRequest()  测试 字符串" + strconv.Itoa(a+b)
}

// 工厂函数
func NewAdaptee() Adaptee {
	return &adapeeImpl{}
}

func main() {
	// 适配器
	ad := NewAdaptee()
	// 传递进入
	ta := NewAdaptor(ad)
	res := ta.Request(2, 3)
	fmt.Println(res)
}
