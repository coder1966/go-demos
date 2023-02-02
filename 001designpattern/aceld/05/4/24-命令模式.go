/*

 */

package main

import "fmt"

// ====== 命令的接受者 核心计算模块
type Docker struct {
}

func (d *Docker) treatEye() {
	fmt.Println("医生 治疗眼睛")
}

func (d *Docker) treatNose() {
	fmt.Println("医生 治疗  鼻子")
}

// 抽象的命令
type Command interface {
	Treat() // 抽象的指令接口
}

// 病单，把病人  核心模块医生 解耦
// 治疗眼睛的病单
type CommandTreatEye struct {
	docker *Docker // 关联医生
}

func (c *CommandTreatEye) Treat() {
	c.docker.treatEye()
}

// 治疗鼻子的病单
type CommandTreatNose struct {
	docker *Docker // 关联医生
}

func (c *CommandTreatNose) Treat() {
	c.docker.treatNose()
}

// ====== 命令的调用者 护士
type Nurse struct {
	CmdList []Command // 病单数组，抽象的
}

// ====== 发送命令 发送病单
func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}

	for _, cmd := range n.CmdList {
		cmd.Treat() // 多态现象，调用具体的命令。就调用 病单 绑定的 医生方法
	}
}

// 病人
func main() {
	// 依赖病单，填写病单 让医生看病
	d := new(Docker) // 只是new 一个对象，没有接口

	// 只是创建病单，等待护士
	cmdEye := CommandTreatEye{d}
	cmdNose := CommandTreatNose{d}

	// 护士
	nurse := new(Nurse)
	// 收集 病单
	nurse.CmdList = append(nurse.CmdList, &cmdEye, &cmdNose)
	// 执行
	nurse.Notify()

}
