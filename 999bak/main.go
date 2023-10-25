package main

import "fmt"

func main() {
	var value float64 = .0999
	s := fmt.Sprintf("%g", value)

	_ = s

	s = s + ""
}
