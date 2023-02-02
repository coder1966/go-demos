package composite

import "fmt"

// 孩子

type Leaf struct {
	component
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

// 多态，重构.否则， main 不打印 叶子
func (l *Leaf) Print(pre string) {
	fmt.Println(pre, l.Name())
}

type Composite struct {
	component
	children []Component // 代表叶子的集合
}

// 创建一个组合结构体
func NewComposite() *Composite {
	return &Composite{children: make([]Component, 0)}
}

// 这里产生了go 的多态
func (c *Composite) AddChild(child Component) {
	child.SetParent(c)                     // 设置父节点
	c.children = append(c.children, child) // 加入叶子|孩子
}
func (c *Composite) Print(pre string) {
	fmt.Println(pre, c.name)
	pre += "   "
	for _, comp := range c.children {
		comp.Print(pre) // 有点递归

	}
}
