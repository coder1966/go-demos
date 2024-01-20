package main

import "math"

func avgFloat(x ...float64) float64 {
	var sum float64
	for _, v := range x {
		sum += v
	}

	if len(x) > 0 {
		return sum / float64(len(x))
	} else {
		return 0
	}
}

func diskUsage(d1, d2, d3, d4 float64) float64 {
	absD1 := math.Abs(d1)
	absD2 := math.Abs(d2)
	absD3 := math.Abs(d3)
	absD4 := math.Abs(d4)

	sum := absD1 + absD2 + absD3
	total := sum + absD4

	if total > 0 {
		return sum / total * 100
	} else {
		return 0
	}
}

func sum(d ...float64) float64 {
	var sum float64
	for _, v := range d {
		sum += math.Abs(v)
	}

	return sum
}
