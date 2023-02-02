package goctxtimeout

import (
	"context"
	"time"

	"godemos/goctxtimeout/dosome"
	"godemos/goctxtimeout/dowithctx"
)

const timeout = 3

func Goctxtimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*timeout)
	defer cancel()
	_, _ = dowithctx.DoWithContext(ctx, func() (interface{}, error) {
		return nil, dosome.DoSome(2)
	})
}
