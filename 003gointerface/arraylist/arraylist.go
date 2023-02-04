package arraylist

import (
	"errors"
	"fmt"
)

// List 接口
type List interface {
	Size() int                                  // 数组大小
	Get(index int) (interface{}, error)         // 抓取第几个元素
	Set(index int, newVal interface{}) error    // 修改数据
	Insert(index int, newVal interface{}) error // 插入数据
	Append(newVal interface{})                  // 追加数据
	Clear()                                     // 清空
	Delete(index int) error                     // 删除
	String() string                             // 返回字符串
	Iterator() Iterator                         // 迭代器接口
}

// ArrayList 数据结构
type ArrayList struct {
	dataStore []interface{} // 数组存储
	theSize   int           // 数组大小
}

func (List *ArrayList) Set(index int, newVal interface{}) error {
	panic("implement me")
}

func (List *ArrayList) Insert(index int, newVal interface{}) error {
	panic("implement me")
}

func (List *ArrayList) Append(newVal interface{}) {
	panic("implement me")
}

func (List *ArrayList) Clear() {
	panic("implement me")
}

func (List *ArrayList) Delete(index int) error {
	panic("implement me")
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)                      // 初始化结构体
	list.dataStore = make([]interface{}, 0, 10) // 开辟空间10个
	list.theSize = 0
	return list

}

func (List *ArrayList) Size() int {
	return List.theSize
}

func (List *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index > List.theSize-1 {
		return nil, errors.New("索引越界")
	}
	return List.dataStore[index], nil
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}
