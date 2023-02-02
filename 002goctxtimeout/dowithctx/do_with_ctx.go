package dowithctx

import (
	"context"
	"fmt"
	"time"
)

func DoWithContext(ctx context.Context, fn func() (interface{}, error)) (interface{}, error) {
	resCh := make(chan interface{})
	errCh := make(chan error)
	go func() {
		res, err := fn()
		if err != nil {
			errCh <- err
		} else {
			// if err, need not res
			resCh <- res
		}
	}()

	for {
		select {
		case <-ctx.Done():
			// 超时返回
			// fmt.Println("timeout failed")
			return nil, fmt.Errorf("timeout failed")
		case res := <-resCh:
			// 正常返回
			return res, nil
		case err := <-errCh:
			// 出错返回
			return nil, err
		case <-time.After(time.Millisecond * 500):
			// 每1秒检查一次是否超时
			continue
		}

	}

}
