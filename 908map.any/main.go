package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

func main() {

	list = make(map[string]Student)

	student := Student{"Aceld"}

	list["student"] = student
	// map[string]Student 的 value 是一个 Student 结构值，所以当list["student"] = student,是一个值拷贝过程。而list["student"]则是一个值引用。那么值引用的特点是只读。所以对list["student"].Name = "LDB"的修改是不允许的。
	// list["student"].Name = "LDB"
	// (*(list["student"])).Name = "LDB"
	// aaa:=list["student"]
	// (*aaa).Name="LDB"
	fmt.Println(list["student"].Name) // 允许

	fmt.Println(list["student"])
}
