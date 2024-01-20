package benchmarktest

import (
	"math/rand"
	"testing"
	"time"
)

/*
$ go test -benchmem -bench .
Benchmark_addSlice-8    1000000000               0.008168 ns/op        0 B/op          0 allocs/op
*/

/*
GODEBUG=gctrace=1 go test -benchmem -bench=.
gc 1 @0.020s 1%: 0.063+4.4+0.038 ms clock, 0.50+0.58/1.0/0.96+0.31 ms cpu, 3->4->1 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 2 @0.031s 1%: 0.075+2.1+0.016 ms clock, 0.60+0.18/1.7/0+0.12 ms cpu, 3->4->1 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 3 @0.037s 2%: 0.054+0.84+0.050 ms clock, 0.43+1.2/1.4/0+0.40 ms cpu, 3->3->1 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 4 @0.041s 2%: 0.024+2.1+0.026 ms clock, 0.19+0.12/1.1/2.0+0.21 ms cpu, 3->4->2 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 5 @0.050s 3%: 0.27+1.3+0.002 ms clock, 2.2+0.72/1.6/0+0.022 ms cpu, 3->4->2 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 6 @0.052s 4%: 0.14+1.8+0.013 ms clock, 1.1+0.62/3.0/0+0.10 ms cpu, 4->5->2 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 7 @0.056s 5%: 0.092+1.5+0.006 ms clock, 0.74+0.79/2.8/0.33+0.053 ms cpu, 4->5->2 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 8 @0.061s 5%: 0.032+0.70+0.005 ms clock, 0.26+0/1.2/1.3+0.042 ms cpu, 3->4->2 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 9 @0.064s 5%: 0.20+1.5+0.007 ms clock, 1.6+1.0/1.7/2.6+0.063 ms cpu, 3->4->2 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 10 @0.066s 6%: 0.034+0.82+0.058 ms clock, 0.27+0.96/1.2/0.26+0.46 ms cpu, 4->5->2 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 11 @0.068s 6%: 0.026+1.3+0.002 ms clock, 0.20+0.15/1.0/1.3+0.019 ms cpu, 4->6->3 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 12 @0.072s 6%: 0.081+0.97+0.010 ms clock, 0.65+0.52/1.7/0+0.085 ms cpu, 6->7->2 MB, 7 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 13 @0.080s 7%: 1.2+0.47+0.008 ms clock, 9.9+0.26/0.88/1.3+0.067 ms cpu, 5->5->2 MB, 6 MB goal, 0 MB stacks, 0 MB globals, 8 P
# github.com/coder1966/go-demonew/a0700unittest/benchmarktest.test
gc 1 @0.002s 6%: 0.047+1.1+0.020 ms clock, 0.38+0.11/1.4/0.62+0.16 ms cpu, 3->4->3 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 2 @0.007s 6%: 0.031+1.3+0.009 ms clock, 0.25+0.41/2.0/1.1+0.078 ms cpu, 6->7->5 MB, 6 MB goal, 0 MB stacks, 0 MB globals, 8 P
# github.com/coder1966/go-demonew/a0700unittest/benchmarktest.test
gc 1 @0.001s 10%: 0.030+1.5+0.010 ms clock, 0.24+0.095/2.1/1.3+0.081 ms cpu, 4->7->6 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 2 @0.005s 9%: 0.011+1.1+0.002 ms clock, 0.088+0.082/2.2/0.15+0.017 ms cpu, 14->14->13 MB, 12 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 3 @0.065s 1%: 0.025+1.3+0.021 ms clock, 0.20+0.10/1.1/1.2+0.17 ms cpu, 22->23->12 MB, 26 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 1 @0.000s 4%: 0.006+0.091+0.005 ms clock, 0.051+0/0.098/0.059+0.046 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 2 @0.000s 6%: 0.011+0.079+0.001 ms clock, 0.090+0/0.12/0.051+0.013 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 3 @0.000s 7%: 0.009+0.063+0.001 ms clock, 0.075+0/0.093/0.067+0.014 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
goos: linux
goarch: amd64
pkg: github.com/coder1966/go-demonew/a0700unittest/benchmarktest
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
Benchmark_addSlice/###Benchmark_addSlice###-8           gc 4 @0.001s 7%: 0.011+0.042+0.001 ms clock, 0.088+0/0.074/0.071+0.015 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 5 @0.001s 7%: 0.012+0.062+0.001 ms clock, 0.099+0/0.083/0.048+0.011 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 6 @0.001s 7%: 0.011+0.056+0.001 ms clock, 0.091+0/0.089/0.083+0.014 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 7 @0.002s 7%: 0.017+0.068+0.002 ms clock, 0.14+0/0.089/0.075+0.016 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 8 @0.002s 7%: 0.011+0.069+0.001 ms clock, 0.089+0/0.093/0.062+0.012 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
1000000000               0.0000814 ns/op               0 B/op          0 allocs/op
*/

// GODEBUG=gctrace=1 go test -benchmem -bench Benchmark_addS
func Benchmark_addS(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AddS()
	}
}

/*
GODEBUG=gctrace=1 go test -benchmem -bench=.
gc 1 @0.016s 2%: 0.070+1.9+0.048 ms clock, 0.56+0.91/1.7/0.66+0.38 ms cpu, 3->3->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 2 @0.028s 6%: 1.1+0.85+0.018 ms clock, 9.5+1.1/1.1/0+0.14 ms cpu, 3->3->1 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 3 @0.033s 6%: 0.10+3.9+0.015 ms clock, 0.85+0.68/1.4/0.24+0.12 ms cpu, 3->4->2 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 4 @0.043s 8%: 0.97+4.6+0.024 ms clock, 7.7+0.40/5.3/0.19+0.19 ms cpu, 3->5->2 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 5 @0.056s 7%: 0.034+0.71+0.010 ms clock, 0.27+0.58/1.2/1.8+0.086 ms cpu, 6->6->1 MB, 6 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 6 @0.058s 7%: 0.077+1.5+0.025 ms clock, 0.62+0.45/2.6/3.0+0.20 ms cpu, 3->4->2 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 7 @0.063s 7%: 0.094+0.94+0.004 ms clock, 0.75+0.12/1.3/1.7+0.033 ms cpu, 4->4->1 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 8 @0.066s 7%: 0.025+0.54+0.015 ms clock, 0.20+0.060/0.85/1.1+0.12 ms cpu, 3->3->1 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 9 @0.067s 8%: 0.039+0.58+0.002 ms clock, 0.31+0.78/1.0/0.27+0.017 ms cpu, 3->4->2 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 10 @0.068s 8%: 0.11+0.91+0.016 ms clock, 0.89+0.57/1.2/0.43+0.13 ms cpu, 3->5->2 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 11 @0.070s 8%: 0.072+0.86+0.009 ms clock, 0.58+0.46/1.3/0.23+0.076 ms cpu, 4->5->2 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 12 @0.072s 8%: 0.031+0.76+0.003 ms clock, 0.25+0.20/1.3/1.2+0.029 ms cpu, 4->5->2 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 13 @0.075s 9%: 0.76+0.85+0.063 ms clock, 6.1+1.0/1.2/0.23+0.50 ms cpu, 4->5->2 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 8 P
# github.com/coder1966/go-demonew/a0700unittest/benchmarktest.test
gc 1 @0.004s 4%: 0.009+5.5+0.008 ms clock, 0.078+0.23/3.1/0.65+0.071 ms cpu, 3->5->3 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 2 @0.013s 5%: 0.042+1.5+0.009 ms clock, 0.34+0.16/2.7/1.2+0.074 ms cpu, 6->7->5 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 14 @0.080s 9%: 0.033+0.58+0.008 ms clock, 0.26+0.28/0.90/1.5+0.066 ms cpu, 4->4->2 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
# github.com/coder1966/go-demonew/a0700unittest/benchmarktest.test
gc 1 @0.000s 9%: 0.011+1.0+0.009 ms clock, 0.092+0.081/1.2/0.077+0.074 ms cpu, 4->7->6 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 2 @0.003s 6%: 0.009+0.61+0.002 ms clock, 0.075+0.096/0.82/0.92+0.019 ms cpu, 14->14->13 MB, 12 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 1 @0.000s 3%: 0.006+0.084+0.001 ms clock, 0.054+0/0.11/0.065+0.015 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 2 @0.000s 4%: 0.006+0.072+0.002 ms clock, 0.055+0/0.099/0.069+0.019 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 3 @0.001s 6%: 0.012+0.070+0.001 ms clock, 0.10+0/0.094/0.057+0.013 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
goos: linux
goarch: amd64
pkg: github.com/coder1966/go-demonew/a0700unittest/benchmarktest
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
Benchmark_addSlice/###Benchmark_addSlice###-8           gc 4 @0.001s 6%: 0.013+0.078+0.002 ms clock, 0.10+0/0.099/0.080+0.016 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 5 @0.001s 6%: 0.011+0.082+0.002 ms clock, 0.088+0/0.10/0.056+0.016 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 6 @0.002s 6%: 0.009+0.064+0.001 ms clock, 0.076+0/0.10/0.094+0.014 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 7 @0.002s 6%: 0.009+0.071+0.001 ms clock, 0.078+0/0.094/0.053+0.014 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 8 @0.002s 6%: 0.011+0.085+0.002 ms clock, 0.094+0/0.096/0.085+0.016 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
1000000000               0.0001177 ns/op               0 B/op          0 allocs/op
gc 9 @0.003s 6%: 0.011+0.062+0.001 ms clock, 0.095+0/0.087/0.048+0.012 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 10 @0.003s 7%: 0.008+0.067+0.001 ms clock, 0.065+0/0.093/0.065+0.012 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
Benchmark_addSlicePool/###Benchmark_addSlicePool###-8   gc 11 @0.003s 6%: 0.010+0.071+0.001 ms clock, 0.082+0/0.099/0.063+0.011 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 12 @0.004s 6%: 0.014+0.061+0.001 ms clock, 0.11+0/0.090/0.079+0.015 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 13 @0.004s 6%: 0.012+0.093+0.001 ms clock, 0.097+0/0.11/0.066+0.012 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 14 @0.005s 6%: 0.013+0.054+0.002 ms clock, 0.10+0/0.089/0.078+0.016 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
gc 15 @0.005s 6%: 0.012+0.067+0.002 ms clock, 0.098+0/0.11/0.062+0.018 ms cpu, 0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P (forced)
1000000000               0.0002540 ns/op               0 B/op          0 allocs/op
*/

// GODEBUG=gctrace=1 go test -benchmem -bench Benchmark_addSPool
func Benchmark_addSPool(b *testing.B) {
	// b.Run("###Benchmark_addSPool###", func(b *testing.B) {
	// 	addSPool()
	// })

	for n := 0; n < b.N; n++ {
		addSPool()
	}
}

// GODEBUG=gctrace=1 go test -benchmem -bench BenchmarkFib
func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}

func BenchmarkParallelFib(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fib(30)
		}
	})
}

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func BenchmarkGenerateWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateWithCap(1000000)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(1000000)
	}
}

// GODEBUG=gctrace=1 go test -benchmem -bench Benchmark_addIOnly
// 838           1369876 ns/op         8192020 B/op       1000 allocs/op
func Benchmark_addIOnly(b *testing.B) {
	for n := 0; n < b.N; n++ {
		addI()
	}
}

// GODEBUG=gctrace=1 go test -benchmem -bench Benchmark_addIPool
func Benchmark_addIPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		addIPool()
	}
}
