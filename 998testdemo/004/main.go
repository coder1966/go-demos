package main

import (
	"fmt"
)

func main2() {
	s1 := []interface{}{"aaa", 1, 2, 3}
	s2 := []interface{}{"bbb", 3, 4, 5}
	v := []interface{}{s1, s2}

	fmt.Println(v)
}
