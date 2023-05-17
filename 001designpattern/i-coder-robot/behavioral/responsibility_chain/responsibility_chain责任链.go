package main

import (
	"fmt"
	"strconv"
)

type Handler interface {
	Handler(handlerID int) string
}

type handler struct {
	name      string
	next      Handler
	handlerID int
}

func NewHandler(n string, next Handler, hID int) *handler {
	return &handler{
		name:      n,
		next:      next,
		handlerID: hID,
	}
}

func (h *handler) Handler(hID int) string {
	// 这个ID是我，我就签字
	if h.handlerID == hID {
		return h.name + " handled " + strconv.Itoa(hID)
	}
	// 我是最后一个签字的，内有后续了
	if h.next == nil {
		return ""
	}
	// 后面还有人，就继续执行
	return h.next.Handler(hID)
}

func main() {
	wang := NewHandler("laowang", nil, 1)
	zhang := NewHandler("laozhang", wang, 2)

	r := wang.Handler(1)
	fmt.Println(r)
	r = zhang.Handler(2)
	fmt.Println(r)
}
