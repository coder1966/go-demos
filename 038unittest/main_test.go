package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
main 单元测试
*/
func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

/*
带参数 普通单元测试
*/
func Test_add(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{
				x: 1,
				y: 3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
	泛型 单元测试
*/

// 必须在测试函数外单独定义测试用例的结构体
type testCase[T int | float64 | string] struct {
	name string
	a    T
	b    T
	want T
}

// 同时还必须在测试函数外定义一个执行泛型用例的泛型函数
func runTestCases[T int | float64 | string](t *testing.T, cases []testCase[T]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := addGenerics(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("xxxxxxxxxxxxxx")
			}
		})
	}
}

// 单元测试函数
func TestAdd(t *testing.T) {
	intTestCases := []testCase[int]{
		{
			name: "ok",
			a:    1,
			b:    1,
			want: 2,
		},
		{
			name: "ok2",
			a:    10,
			b:    10,
			want: 20,
		},
	}
	strCases := []testCase[string]{
		{
			name: "ok",
			a:    "A",
			b:    "B",
			want: "AB",
		},
		{
			name: "ok2",
			a:    "Hello",
			b:    "World",
			want: "HelloWorld",
		},
	}
	runTestCases(t, intTestCases)
	runTestCases(t, strCases)
}

/*
带返回 error 单元测试
*/
func Test_wantLess5(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "ok 1",
			args: args{
				i: 1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "ok 5",
			args: args{
				i: 5,
			},
			want:    1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := wantLess5(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("wantLess5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			assert.Equal(t, got, tt.want)

			if got != tt.want {
				t.Errorf("wantLess5() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
Benchmark 测试

$ go test -bench .
goos: linux
goarch: amd64
pkg: godemos/700unittest
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
Benchmark_add-8         1000000000               0.2559 ns/op
PASS
ok      godemos/700unittest     0.293s
*/
func Benchmark_add(b *testing.B) {
	for n := 0; n < b.N; n++ {
		add(30, n)
	}
}
