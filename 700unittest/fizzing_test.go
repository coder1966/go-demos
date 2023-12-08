package main

import (
	comparefloat "godemos/711comparefloat"
	"testing"
)

// func Fuzz_avgFloat(f *testing.F) {

// 	f.Fuzz(func(t *testing.T) {
// 		if got := avgFloat(tt.args.x, tt.args.y); got != tt.want {
// 			t.Errorf("addFuzz() = %v, want %v", got, tt.want)
// 		}
// 	})

// }

// func Fuzz_avgFloat2(f *testing.F) {
// 	f.Fuzz(func(t *testing.T, orig1 float64, orig2 float64) {
// 		got := avgFloat(orig1, orig2)
// 		want := avgFloat(orig1, orig2)

// 		if got != want {
// 			t.Errorf("got: %f, want: %f; orig: %f, %f ;", got, want, orig1, orig2)
// 		}

// 	})
// }

// go test -fuzz=Fuzz -fuzztime 5s
func Fuzz_diskUsage(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1, d2, d3, d4 float64) {
		got := diskUsage(d1, d2, d3, d4)

		want := 0.0
		if sum(d1, d2, d3, d4) > 0 {
			want = sum(d1, d2, d3) / sum(d1, d2, d3, d4) * 100
		}

		if !comparefloat.CompareFloat(got, want) {
			t.Errorf("got: %v, want: %v; orig: %v, %v, %v, %v;", got, want, d1, d2, d3, d4)
		}

		// if !comparefloat.CompareFloatOld(got, want) {
		// 	t.Errorf("got: %v, want: %v; orig: %v, %v, %v, %v;", got, want, d1, d2, d3, d4)
		// }

	})
}

// func Test_diskUsage(t *testing.T) {
// 	type args struct {
// 		d1    float64
// 		d2    float64
// 		d3    float64
// 		total float64
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want float64
// 	}{
// 		{
// 			name: "string{}",
// 			args: args{
// 				d1:    1,
// 				d2:    2,
// 				d3:    3,
// 				total: 10,
// 			},
// 			want: 0.6,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := diskUsage(tt.args.d1, tt.args.d2, tt.args.d3, tt.args.total); got != tt.want {
// 				t.Errorf("diskUsage() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
