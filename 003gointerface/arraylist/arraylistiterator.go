package arraylist

// Iterator 第一一个接口
type Iterator interface {
	HasNext() bool              // 是否有下一个
	Next() (interface{}, error) // 下一个
	Remove()                    // 删除
	GetIndex()                  // 得到当前的索引
}

// Iterable 第二个个 接口；？接口嵌套？
type Iterable interface {
	Iterator() Iterator // 构造初始化接口
}

// ArrayListIterator 构造指针访问数组
type ArrayListIterator struct {
	list         *ArrayList // 数组指针(一个类包含其他的类)
	currentIndex int        // 当前索引
}

func (a ArrayListIterator) HasNext() bool {
	panic("implement me")
}

func (a ArrayListIterator) Next() (interface{}, error) {
	panic("implement me")
}

func (a ArrayListIterator) Remove() {
	a.currentIndex--
	a.list.Delete(a.currentIndex)
}

func (a ArrayListIterator) GetIndex() {
	panic("implement me")
}

// Iterator 把 type Iterable interface 这个接口实现一下
func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator) // it代替这个 Iterator；构造迭代器
	it.currentIndex = 0
	it.list = list
	return it // 必须把“第一个接口”的4个函数实现才行。否则飘红

}
