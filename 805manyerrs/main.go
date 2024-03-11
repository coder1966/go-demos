package main

// https://www.bilibili.com/video/BV1sM4y1o7Rt/?spm_id_from=pageDriver&vd_source=551c7981013130e9e2f1594d47bd7ca0
// 解除不断调用 if err！=nil

// err 发到结构体里面，返回结构体，链式调用
type Book struct {
	Name   string
	Price  int
	Store  int
	Member int
	err    error
}

func clientExample() {
	book2 := &Book{}
	book2.CalcDiscount(99).
		IfSale()
	//  ...
}

func (b *Book) CalcDiscount(count int) *Book {
	// 其实是检测上一个执行是否有错误
	if b.err != nil {
		return b
	}
	// 业务 ...
	// b.err= 如果有错误

	return b
}

func (b *Book) IfSale() *Book {
	// 其实是检测上一个执行是否有错误
	if b.err != nil {
		return b
	}
	// 业务 ...
	// b.err= 如果有错误

	return b
}

func main() {

}
