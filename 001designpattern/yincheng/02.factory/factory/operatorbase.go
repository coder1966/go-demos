package factory

// 数据
type OperatorBase struct {
	left, right int
}

// 赋值
func (o *OperatorBase) SetLeft(data int) {
	o.left = data
}
func (o *OperatorBase) SetRight(data int) {
	o.right = data
}
