package main

import "fmt"

func main() {
	var value float64 = .0999
	s := fmt.Sprintf("%g", value)

	a1 := f1()
	a2 := f2()
	a3 := f3()

	_, _, _ = a1, a2, a3
	_ = s

	s = s + ""
}
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
