package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	mock 一个 interface 单元测试

	要点在：
	type Collector interface {
		GetTime() int64
		Stop()
	}

	要点在：feeder 可以取回来 Feed 的数据
	type mockFeed struct{ ts int64 }
*/

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
			mockFeed := &mockFeed{}
			i := &input{
				collector: mockCollect,
				feeder:    mockFeed,
			}

			mockCollect.setTimStamp(tt.fields.ts)

			i.doGetter()

			got := mockFeed.getTsBack()

			assert.Equal(t, got, tt.want)

		})
	}
}

// mock
type mockCollect struct{ ts int64 }

func (c *mockCollect) GetTime() int64      { return c.ts }
func (c *mockCollect) Stop()               {}
func (c *mockCollect) setTimStamp(i int64) { c.ts = i }

// mock
type mockFeed struct{ ts int64 }

func (f *mockFeed) Feed(ts int64) error {
	f.ts = ts
	// fmt.Println("Yes, mock feeding: ", ts)
	return nil
}

func (f *mockFeed) Clear() {}
func (f *mockFeed) getTsBack() int64 {
	return f.ts
}
