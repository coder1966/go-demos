package main

import (
	"testing"
)

func Test_input_doGetter(t *testing.T) {
	type fields struct {
		ts int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "ok 12345",
			fields: fields{
				ts: 12345,
			},
			want: 12345,
		},
		{
			name: "ok 67890",
			fields: fields{
				ts: 67890,
			},
			want: 67890,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCollect := &mockCollect{}
			i := &input{
				collector: mockCollect,
			}

			mockCollect.setTimStamp(tt.fields.ts)

			if got := i.doGetter(); got != tt.want {
				t.Errorf("input.doGetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockCollect struct{ ts int64 }

func (c *mockCollect) GetTime() int64      { return c.ts }
func (c *mockCollect) Stop()               {}
func (c *mockCollect) setTimStamp(i int64) { c.ts = i }
