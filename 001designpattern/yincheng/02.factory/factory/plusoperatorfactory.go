package factory

// 操作的抽象
type PlusOperatorFactory struct{}

// 操作类，依赖数据。所以引入数据
type PlusOperator struct {
	*OperatorBase
}

// 对刚才数据做实际的操作
func (o *PlusOperator) Result() int {
	return o.left + o.right
}

func (PlusOperatorFactory) Creat() Operator {
	return &PlusOperator{OperatorBase: &OperatorBase{}}
}
