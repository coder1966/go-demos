package main

import (
	"fmt"
)

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3},
}

func main() {
	n := make(map[string]Math)
	n["foo"] = Math{2, 3}
	nn := n["foo"]
	nn.x = 4
	// n["foo"].x = 4

	// m["foo"].x = 4
	fmt.Println(m["foo"].x)
}
