package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	mock 一个 func 单元测试

	要点在： type getTime func() int64
*/

func Test_input_doGetter(t *testing.T) {
	type fields struct {
		getter getTime
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "ok 12345",
			fields: fields{
				getter: getTimeMock_12345,
			},
			want: 12345,
		},
		{
			name: "ok 67890",
			fields: fields{
				getter: getTimeMock_67890,
			},
			want: 67890,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &input{
				getter: tt.fields.getter,
			}

			got := i.doGetter()

			assert.Equal(t, got, tt.want)
		})
	}
}

func getTimeMock_12345() int64 {
	return 12345
}

func getTimeMock_67890() int64 {
	return 67890
}
