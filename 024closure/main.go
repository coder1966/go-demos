package main

import (
	"fmt"
	"strings"
)

// 函数，返回值是一个函数 func(int) int
func adder() func(int) int {
	var x int
	return func(y int) int {
		// 返回的函数，是一个闭包
		x += y
		return x
	}
}

// 闭包进阶示例1：
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包进阶示例2：
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 闭包进阶示例3：
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	// 变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。 在f的生命周期内，变量x也一直有效。
	var f = adder()

	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	// 新的实例
	f1 := adder()
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90

	// 闭包进阶示例1：
	var ff = adder2(10)
	fmt.Println(ff(10)) //20
	fmt.Println(ff(20)) //40
	fmt.Println(ff(30)) //70

	ff1 := adder2(20)
	fmt.Println(ff1(40)) //60
	fmt.Println(ff1(50)) //110

	// 闭包进阶示例2：
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt

	// 闭包进阶示例3：
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7

}
