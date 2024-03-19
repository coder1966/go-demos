package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Man2 struct {
	Name   string
	Age    int
	Detail []byte
}

func main() {

	func2()
}

func func2() {
	li4 := &Man2{Name: "li-si", Age: 18, Detail: []byte("Hello World")}
	// encode ============
	var b bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&b) // Will write to network.

	err := enc.Encode(li4)
	if err != nil {
		fmt.Println("ERR: ", err.Error())
	} else {
		fmt.Println("GOT: ", string(b.String()))
	}

	// decode ===========
	var duplicate Man2

	// var network bytes.Buffer  // Stand-in for the network connection
	dec := gob.NewDecoder(&b) // Will read from network.

	// 假设 network 已经被编码数据填充
	// ...

	err = dec.Decode(&duplicate)
	if err != nil {
		fmt.Println("ERR: ", err.Error())
	}
	fmt.Println("GOT: ", duplicate)
}
