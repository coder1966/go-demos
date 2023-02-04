package gointerface

import (
	"fmt"

	"godemos/gointerface/stackarray"
)

// func main6() {

// 	// 定义接口对象，赋值对象必须实现接口的所有方法
// 	var list arraylist.List = arraylist.NewArrayList()
// 	list.Append("A1")

// 	for it := list.Iterator(); it.HasNext(); {
// 		item, _ := it.Next()
// 		fmt.Println(item)
// 	}

// }

func Interfacemain() {

	myStack := stackarray.NewStack()
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)

	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
}
