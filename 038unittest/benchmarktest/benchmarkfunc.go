package benchmarktest

import (
	"strings"
	"sync"
)

const (
	m = 1000
	n = 1000000
)

var mockS1 = strings.Repeat("a", n)
var mockS2 = strings.Repeat("b", n)

func AddS() {
	for i := 0; i < m; i++ {
		var s string
		if i%2 == 0 {
			s = mockS1
		} else {
			s = mockS2
		}

		_ = s
	}
}

var reqPool = sync.Pool{
	New: func() interface{} {
		return new(string)
	},
}

func addSPool() {
	for i := 0; i < m; i++ {
		s := reqPool.Get().(*string)
		if i%2 == 0 {
			*s = mockS1
		} else {
			*s = mockS2
		}
		_ = s
	}
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

var l = 1000

var iPool = sync.Pool{
	New: func() interface{} {
		return new([]int)
	},
}

func addIPool() {
	s := iPool.Get().(*[]int)

	for i := 0; i < l; i++ {
		(*s) = append((*s), i)
	}
	_ = s
}

func addI() {
	s := make([]int, 0, 0)

	for i := 0; i < l; i++ {
		s = append(s, i)
	}
	_ = s
}
