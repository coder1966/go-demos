package main

import "fmt"

type Docker struct {
}

func (d *Docker) treatEye() {
	fmt.Println("医生 治疗眼睛")
}

func (d *Docker) treatNose() {
	fmt.Println("医生 治疗  鼻子")
}

// 病人
func main() {
	d := new(Docker)
	d.treatEye()
	d.treatNose()
}
