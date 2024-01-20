package a0037cassa

import "testing"

func Test_doCass(t *testing.T) {

	tests := []struct {
		name string
	}{
		{name: "192.168.56.15"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doCass(tt.name)
		})
	}
}
