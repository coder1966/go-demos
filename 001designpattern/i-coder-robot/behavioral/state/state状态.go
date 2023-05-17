package main

import "fmt"

type State interface {
	On(m *Machine)
	Off(m *Machine)
}
type Machine struct {
	current State
}

// 状态机 维护了  ON struct{}  OFF struct{}
func NewMachine() *Machine            { return &Machine{NewOFF()} } // 新建状态机，默认off
func (m *Machine) setCurrent(s State) { m.current = s }
func (m *Machine) On()                { m.current.On(m) }
func (m *Machine) Off()               { m.current.Off(m) }

type ON struct{}

func NewON() State          { return &ON{} }
func (o *ON) On(m *Machine) { fmt.Println("设备已经是开启状态了。。。") }
func (o *ON) Off(m *Machine) {
	fmt.Println("从on状态转变成off了。。。")
	m.setCurrent(NewOFF())
}

type OFF struct{}

func NewOFF() State { return &OFF{} }
func (o *OFF) On(m *Machine) {
	fmt.Println("从off状态转变成on了。。。")
	m.setCurrent(NewON())
}
func (o *OFF) Off(m *Machine) { fmt.Println("已经是off了。。。") }

func main() {
	m := NewMachine()
	m.Off()
	m.On()
	m.On()
	m.Off()
}
