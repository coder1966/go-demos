package main

import (
	"encoding/json"
	"fmt"
)

type Man struct {
	Name string
	Age  int
}

type Man2 struct {
	Name   string
	Age    int
	Detail []byte
}

func main() {
	func1()
	func2()
}

func func1() {
	li4 := &Man{Name: "li-si", Age: 18}
	// encode ============
	b, err := json.Marshal(li4)
	if err != nil {
		fmt.Println("ERR: ", err.Error())
	} else {
		fmt.Println("GOT: ", string(b))
	}

	// decode ===========
	var duplicate Man
	err = json.Unmarshal(b, &duplicate)
	if err != nil {
		fmt.Println("ERR: ", err.Error())
	}
	fmt.Println("GOT: ", duplicate)
}

func func2() {
	li4 := &Man2{Name: "li-si", Age: 18, Detail: []byte("Hello World")}
	// encode ============
	b, err := json.Marshal(li4)
	if err != nil {
		fmt.Println("ERR: ", err.Error())
	} else {
		fmt.Println("GOT: ", string(b))
	}

	// decode ===========
	var duplicate Man2
	err = json.Unmarshal(b, &duplicate)
	if err != nil {
		fmt.Println("ERR: ", err.Error())
	}
	fmt.Println("GOT: ", duplicate)
}
