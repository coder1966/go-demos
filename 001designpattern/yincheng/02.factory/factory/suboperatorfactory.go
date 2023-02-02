package factory

// 操作的抽象
type SubOperatorFactory struct{}

// 操作类，依赖数据。所以引入数据
type SubOperator struct {
	*OperatorBase
}

// 对刚才数据做实际的操作
func (o *SubOperator) Result() int {
	return o.left - o.right
}

func (SubOperatorFactory) Creat() Operator {
	return &SubOperator{OperatorBase: &OperatorBase{}}
}
