package main

import (
	"fmt"
	"time"
)

const times = 1000000
const want = "i"

func main() {
	strs := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
	}

	start := time.Now()
	for i := 0; i < times; i++ {
		for _, v := range strs {
			if v == want {
				break
			}
		}
	}
	fmt.Println("方案1耗时：", time.Since(start))

	start = time.Now()
	for i := 0; i < times; i++ {
		j := 0
		for ; j < len(strs); j++ {

			if strs[j] == want {
				break
			}
		}
	}
	fmt.Println("方案2耗时：", time.Since(start))

	start = time.Now()
	for i := 0; i < times; i++ {
		_ = find(strs, want)
	}
	fmt.Println("方案3耗时：", time.Since(start))
}

func find(strs []string, want string) int {
	for j := 0; j < len(strs); j++ {

		if strs[j] == want {
			return j
		}
	}
	return -1
}
