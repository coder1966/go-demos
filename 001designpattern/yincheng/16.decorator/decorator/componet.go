package decorator

// interface 代表行为
type Component interface {
	Calc() int
}

// struct 代表数据
// 构造完接口，再构造结构体 Concreate创建
type ConcreateComponent struct{}

// 数据和行为绑定一下，实现接口
func (c *ConcreateComponent) Calc() int {
	return 0
}
