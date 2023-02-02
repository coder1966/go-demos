package main

import (
	"fmt"

	"godemos/001designpattern/yincheng/16.decorator/decorator"
)

/*
componet.go 写好 函数，返回 0
add|mul 写一层包装，实现加法|乘法
*/
func main() {
	// 构建
	var c decorator.Component = &decorator.ConcreateComponent{}

	// 加法包装
	c = decorator.WarpAddComponent(c, 10)
	fmt.Println(c.Calc())

	// 乘法包装
	c = decorator.WarpMulComponent(c, 8)
	fmt.Println(c.Calc())
}
