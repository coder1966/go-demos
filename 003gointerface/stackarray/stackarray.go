package stackarray

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type Stack struct {
	dataSource  []interface{}
	capSize     int // 容量/最大大小
	currentSize int // 实际使用大小
}

func NewStack() *Stack {
	myStack := new(Stack)
	myStack.dataSource = make([]interface{}, 0, 10)
	myStack.capSize = 10
	myStack.currentSize = 0
	return myStack
}

func (s *Stack) Clear() {
	s.dataSource = make([]interface{}, 0, 10)
	s.capSize = 10
	s.currentSize = 0
}
func (s *Stack) Size() int {
	return s.currentSize
}
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	last := s.dataSource[s.currentSize-1]
	s.dataSource = s.dataSource[:s.currentSize-1] // 删掉尾巴
	s.currentSize--
	return last
}
func (s *Stack) Push(data interface{}) {
	if !s.IsFull() {
		s.dataSource = append(s.dataSource, data)
		s.currentSize++
	}
}
func (s *Stack) IsFull() bool {
	if s.currentSize >= s.capSize {
		return true
	} else {
		return false
	}
}
func (s *Stack) IsEmpty() bool {
	if s.currentSize == 0 {
		return true
	} else {
		return false
	}
}
