package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("===== start main =====")
	ipt := &input{
		getter: getTimeReal,
	}

	fmt.Println("now ts == ", ipt.doGetter())

	fmt.Println("===== end main =====")
}

func (i *input) doGetter() int64 {
	return i.getter()
}

type input struct {
	getter getTime
}

type getTime func() int64

func getTimeReal() int64 {
	return time.Now().Unix()
}
