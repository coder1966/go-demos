package dowithctx

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestDoWithContext(t *testing.T) {
	type args struct {
		delay int
		fn    func() (interface{}, error)
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "01 toimeout",
			args: args{
				delay: 1,
				fn:    funcNoErr,
			},
			want:    999,
			wantErr: true,
		},
		{
			name: "02 no timeout had err",
			args: args{
				delay: 3,
				fn:    funcWithErr,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "03 no timeout no err",
			args: args{
				delay: 3,
				fn:    funcNoErr,
			},
			want:    999,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(tt.args.delay))
			_ = cancel
			got, err := DoWithContext(ctx, tt.args.fn)
			if (err != nil) && tt.wantErr {
				// 预计是 err
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("DoWithContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoWithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func funcWithErr() (interface{}, error) {
	time.Sleep(time.Second * 2)
	return nil, fmt.Errorf("程序带来的error")
}
func funcNoErr() (interface{}, error) {
	time.Sleep(time.Second * 2)
	return 999, nil
}
