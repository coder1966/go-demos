package a0037cassandra02

import "testing"

func Test_mainCassa(t *testing.T) {
	tests := []struct {
		name string
	}{
		// {name: "172.19.0.2"},
		{name: "192.168.56.15"},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mainCassa(tt.name)
		})
	}
}
