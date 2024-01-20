package main

import (
	"fmt"
	"testing"
)

func main() {

}

func paramMap(m map[string]string) {
	m["c"] = "cc"
	m["d"] = "dd"
}

func paramSlice(s []string) {
	s = append(s, "cc", "dd")
}

func paramSliceV2(s *[]string) {
	*s = append((*s), "cc", "dd")
}

func Test_paramMap(t *testing.T) {

	tests := []struct {
		name string
		m    map[string]string
	}{
		{
			name: "string",
			m:    map[string]string{"a": "aa", "b": "bb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			paramMap(tt.m)
			t.Log(tt.m)
			fmt.Println(tt.m)
		})
	}
}

func Test_paramSlice(t *testing.T) {

	tests := []struct {
		name string
		m    []string
	}{
		{
			name: "string",
			m:    []string{"aa", "bb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			paramSlice(tt.m)
			t.Log(tt.m)
			fmt.Println(tt.m)
		})
	}
}

func Test_paramSliceV2(t *testing.T) {

	tests := []struct {
		name string
		m    []string
	}{
		{
			name: "string",
			m:    []string{"aa", "bb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			paramSliceV2(&tt.m)
			t.Log(tt.m)
			fmt.Println(tt.m)
		})
	}
}
