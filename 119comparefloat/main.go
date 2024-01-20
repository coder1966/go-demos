package main

import "math"

func main() {

}

func compareFloat(a, b float64) bool {
	deviation := 0.0001 // 允许的误差
	if a == 0 && b == 0 {
		return true
	}

	a = math.Abs(a)
	b = math.Abs(b)

	diff := a - b
	max := a
	if a < b {
		diff = b - a
		max = b
	}

	div := diff / max

	return div < deviation
}
