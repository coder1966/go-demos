package a0037cassandra

import (
	"testing"
)

func Test_read(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "ok"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			read()
		})
	}
}

func Test_demo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "ok"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			demo()
		})
	}
}
