package main

import (
	"fmt"

	"godemos/703interfacetest_02/collect"
	"godemos/703interfacetest_02/feed"
)

func main() {
	fmt.Println("===== start main =====")
	ipt := &input{
		collector: collect.NewCollect(),
		feeder:    feed.NewFeed(),
	}

	fmt.Println("error == ", ipt.doGetter())

	fmt.Println("===== end main =====")
}

type input struct {
	collector collect.Collector
	feeder    feed.Feeder
}

func (i *input) doGetter() error {
	ts := i.collector.GetTime()

	i.feeder.Feed(ts)

	return nil
}
