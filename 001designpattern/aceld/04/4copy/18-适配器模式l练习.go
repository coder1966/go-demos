package main

import "fmt"

// 适配的目标 抽象的技能
type Attack interface {
	Fight()
}

// 具体的技能
type Dabaojian struct {
}

func (d *Dabaojian) Fight() {
	fmt.Println("使用了 大宝剑 技能，击杀敌人")
}

// =============
type Hero struct {
	Name   string
	attack Attack // 具体攻击方式

}

func (h *Hero) Skill() {
	fmt.Println(h.Name, "使用了技能")
	h.attack.Fight() // 使用具体的战斗方式
}

// 适配者，使用大保健，改为关机
type PowerOff struct{}

func (p *PowerOff) ShutDown() {
	fmt.Println("电脑即将关机")
}

// 适配器
type Adapter struct {
	PowerOff *PowerOff // 被适配的
}

// 重写Fight
func (a *Adapter) Fight() {
	a.PowerOff.ShutDown()
}
func NewAdaptor(p *PowerOff) *Adapter {
	return &Adapter{p}
}

// ====== 业务逻辑层 ==========
func main() {
	gailunL := Hero{Name: "盖伦", attack: new(Dabaojian)}
	gailunL.Skill()

	// 一使用 大宝剑 就关机
	songjiang := Hero{Name: "宋江", attack: NewAdaptor(new(PowerOff))}
	songjiang.Skill()

}
