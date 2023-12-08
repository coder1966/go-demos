package comparefloat

import (
	"fmt"
	"math"
)

const Min = 0.000001

func CompareFloatOld(x, y float64) bool {
	return math.Abs(x-y) < Min
}

func CompareFloat(x, y float64) bool {
	absX := math.Abs(x)
	absY := math.Abs(y)
	absZ := math.Abs(x - y)

	absMax := absY
	if absY < absX {
		absMax = absX
	}

	if absMax == 0 {
		return true
	}

	fmt.Println(absZ, absMax)

	return absZ/absMax < Min
}
