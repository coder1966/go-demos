/*

 */

package main

import "fmt"

// 核心计算模块
type Docker struct {
}

func (d *Docker) treatEye() {
	fmt.Println("医生 治疗眼睛")
}

func (d *Docker) treatNose() {
	fmt.Println("医生 治疗  鼻子")
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

// 病人
func main() {
	// 依赖病单，填写病单 让医生看病
	d := new(Docker) // 只是new 一个对象，没有接口

	cmdEye := CommandTreatEye{d}
	cmdEye.Treat()

	cmdNose := CommandTreatNose{d}
	cmdNose.Treat()
}
