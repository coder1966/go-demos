package decorator

// 在这里，把 component.go 里面返回 0 重新实现一下

type MulComponent struct {
	// c   Component // 内嵌原来组件。装饰，就要把原来的组件放进来
	Component // 内嵌原来组件。装饰，就要把原来的组件放进来
	num       int
}

// 包装上面的，
func WarpMulComponent(c Component, num int) Component {
	return &MulComponent{c, num}
}

func (c *MulComponent) Calc() int {
	return c.Component.Calc() * c.num // 相当于实现了乘法
}
