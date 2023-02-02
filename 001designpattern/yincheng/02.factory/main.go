package main

import (
	"fmt"

	"godemos/001designpattern/yincheng/02.factory/factory"
)

// //  A X B 运算

func main() {
	var fac factory.OperatorFactory
	fac = factory.PlusOperatorFactory{}
	op := fac.Creat()
	op.SetLeft(50)
	op.SetRight(20)
	fmt.Println(op.Result())

	fac = factory.SubOperatorFactory{}
	op = fac.Creat()
	op.SetLeft(50)
	op.SetRight(20)
	fmt.Println(op.Result())

}
