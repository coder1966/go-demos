package main

import (
	"fmt"

	"godemos/go-strudy/001designpattern/yincheng/06.bulder/builder"
)

func main() {
	bu := &builder.IntBuilder{}
	dic := builder.NewDirector(bu)
	fmt.Println(bu.GetResult())

	dic.MakeData()
	fmt.Println(bu.GetResult())

	fmt.Println("===============")

	bu2 := &builder.StringBuilder{}
	dic2 := builder.NewDirector(bu2)
	fmt.Println(bu2.GetResult())

	dic2.MakeData()
	fmt.Println(bu2.GetResult())
}
