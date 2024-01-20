package main

import "fmt"

func main() {
	fmt.Println("===== start main =====")
	fmt.Println(add(1, 3))
	fmt.Println(addGenerics(1, 3))
	fmt.Println(addGenerics(1.11, 3.33))
	fmt.Println(addGenerics("1.11", "3.33"))

	_, _ = wantLess5(5)

	fmt.Println("===== end main =====")
}

func add(x, y int) int {
	return x + y
}

func addGenerics[T int | float64 | string](x, y T) T {
	return x + y
}

func wantLess5(i int) (int, error) {
	if i < 5 {
		return i, nil

	}
	return i, fmt.Errorf("%d is >= 5", i)
}
