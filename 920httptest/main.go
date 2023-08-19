// https://blog.csdn.net/xixihahalelehehe/article/details/112371200md
package main

import (
	"fmt"
	"regexp"
)

func main() {
	expStr := `\d(+` // 会抛出异常
	// expStr := `^bc`
	strs := []string{
		"bc",
		"abc",
		"abcd",
		"bcd",
	}
	for _, v := range strs {
		regGood(v, expStr)
		// reg(v, expStr)
	}

}

/*
regexp.Compile
// Compile 用来解析正则表达式 expr 是否合法，如果合法，则返回一个 Regexp 对象
// Regexp 对象可以在任意文本上执行需要的操作
func Compile(expr string) (*Regexp, error)
返回一个实现了regexp的对象指针，可以使用返回值调用regexp中定义的方法，如Match，MatchString，find等。

MustCompile 的作用和 Compile 一样
// 不同的是，当正则表达式 str 不合法时，MustCompile 会抛出异常
// 而 Compile 仅返回一个 error 值
func MustCompile(str string) *Regexp

func main() {
reg := regexp.MustCompile(`\w+`)
fmt.Println(reg.FindString("Hello World!"))
// Hello
}

expName,err := regexp.Compile(expStr)
if err != nil {
	l.Errorf("parsing regexp:: %v", err)
	return 0, "", err
}
*/
func reg(str, expStr string) {
	expName := regexp.MustCompile(expStr)
	if expName.MatchString(str) {
		fmt.Println("OK", str, expStr)
	} else {
		fmt.Println("NO", str, expStr)
	}

}
func regGood(str, expStr string) {
	expName, err := regexp.Compile(expStr)
	if err != nil {
		fmt.Println("good error: ", str, expStr, err)
		return
	}
	if expName.MatchString(str) {
		fmt.Println("good OK", str, expStr)
	} else {
		fmt.Println("good NO", str, expStr)
	}

}
