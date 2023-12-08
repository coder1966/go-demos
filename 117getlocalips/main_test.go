package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_testIP(t *testing.T) {

	tests := []struct {
		name   string
		remote string
		want   string
	}{
		{name: "string", remote: "1.2.3.4", want: "1.2.3.4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testIP(tt.remote)

			assert.Equal(t, tt.want, got)

		})
	}
}
