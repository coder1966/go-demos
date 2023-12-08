package main

import (
	"fmt"
	"godemos/704interfacetest_03/collect"
)

func main() {
	fmt.Println("===== start main =====")
	ipt := &input{
		collector: collect.NewCollect(),
	}

	fmt.Println("now ts == ", ipt.doGetter())

	fmt.Println("===== end main =====")
}

func (i *input) doGetter() int64 {
	return i.collector.GetTime()
}

type input struct {
	collector collect.Collector
}
