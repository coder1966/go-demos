package main

import "fmt"

// 结构体：统一对外的接口.测试
type API interface {
	Test() string
}

func NewAPI() API {
	return &APICall{NewAmoudleAPI(), NewBmoudleAPI()} // 返回的是 把两个外观组合起来，统一方法
}

// 把两个外观组合起来，统一方法
type APICall struct {
	// 两种接口
	a AmoudleAPI
	b BmoudleAPI
}

func (api *APICall) Test() string {
	return api.a.TestA() + api.b.TestB()
}

// AAAAAA
type AmoudleAPI interface{ TestA() string } // A 主网
type aMoudleImpl struct{}                   // 接口匹配类
func (api *aMoudleImpl) TestA() string      { return "主网开始运动" }
func NewAmoudleAPI() AmoudleAPI             { return &aMoudleImpl{} }

// BBBBBB
type BmoudleAPI interface{ TestB() string } // B 测试网
type bMoudleImpl struct{}                   // 接口匹配类
func (api *bMoudleImpl) TestB() string      { return "测试网开始运动" }
func NewBmoudleAPI() BmoudleAPI             { return &bMoudleImpl{} }

func main() {
	api := NewAPI()
	fmt.Println(api.Test())
}
