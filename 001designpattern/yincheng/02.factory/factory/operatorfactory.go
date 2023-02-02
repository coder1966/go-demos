package factory

//  A X B 运算

// 实际运行类的接口
type Operator interface {
	SetLeft(int)
	SetRight(int)
	Result() int
}

// 工厂接口。######【返回的是一个接口】######
type OperatorFactory interface {
	Creat() Operator
}
