package main

// https://www.bilibili.com/video/BV19e4y1A7DD/?spm_id_from=333.788&vd_source=551c7981013130e9e2f1594d47bd7ca0

type someInterface interface {
	Do01()
	Do02()
}

type someStruct struct {
}

// Do01 implements someInterface.
func (*someStruct) Do01() {
	panic("unimplemented")
}

// Do02 implements someInterface.
func (*someStruct) Do02() {
	panic("unimplemented")
}

// 选快速修复 。。。。
var _ someInterface = (*someStruct)(nil)

func main() {

}
