package composite

// 合成 混合 模式
// 拆分成不可拆分的最小单元
// 家庭类 父母类 孩子类。

// 一个书，这个有时候是父节点，有时候是叶子节点
type Component interface {
	Parent() Component // 返回 本身
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

const (
	LeafNode      = iota // 叶子节点
	CompositeNode        // 组合节点
)

type component struct {
	parent Component
	name   string
}

// 用它不断地构建
func NewComponent(kind int, name string) Component {
	var c Component
	switch kind { // 根据不同类型，返回不同
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}
	c.SetName(name)
	return c
}

func (c *component) Parent() Component          { return c.parent }
func (c *component) SetParent(parent Component) { c.parent = parent }
func (c *component) Name() string               { return c.name }
func (c *component) SetName(name string)        { c.name = name }

func (c *component) AddChild(Component) {} // 没有实现，通过组合，让child 实现
func (c *component) Print(str string)   {} // 没有实现，通过组合，让child 实现
