package main

import "testing"

func Test_compareFloat(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want bool
	}{
		// {name: "string", a: 0.1, b: 0.1, want: true},
		{name: "string", a: getF(), b: getF(), want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareFloat(tt.a, tt.b); got != tt.want {
				t.Errorf("compareFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getF() float64 {
	f1 := 1.234567892345678923456789
	f2 := 1.123456782345678923456789
	f3 := 1.345678902345678923456789
	f := f1 + f2 + f3
	f = f / f1
	return f

}
