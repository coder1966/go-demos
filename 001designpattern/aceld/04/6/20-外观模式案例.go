package main

import "fmt"

// 电司机
type TV struct{}

func (t *TV) On() {
	fmt.Println("打开 电视机")
}
func (t *TV) Off() {
	fmt.Println("关 电视机")
}

// 音箱
type VoiceBox struct{}

func (v *VoiceBox) On() {
	fmt.Println("打开 音箱")
}
func (v *VoiceBox) Off() {
	fmt.Println("关 音箱")
}

// 灯光
type Light struct{}

func (v *Light) On() {
	fmt.Println("打开 灯光")
}
func (v *Light) Off() {
	fmt.Println("关 灯光")
}

// 游戏机
type Xbox struct{}

func (v *Xbox) On() {
	fmt.Println("打开 游戏机")
}
func (v *Xbox) Off() {
	fmt.Println("关 游戏机")
}

// 麦克风
type MicroPhone struct{}

func (v *MicroPhone) On() {
	fmt.Println("打开 麦克风")
}
func (v *MicroPhone) Off() {
	fmt.Println("关 麦克风")
}

// 投影仪
type Projector struct{}

func (v *Projector) On() {
	fmt.Println("打开 投影仪")
}
func (v *Projector) Off() {
	fmt.Println("关 投影仪")
}

// ------ 外观 家庭影院
type HomePlayerFacade struct {
	tv    TV
	vb    VoiceBox
	Light Light
	xbox  Xbox
	mp    MicroPhone
	pro   Projector
}

//   KTV 模式 组合
func (hp *HomePlayerFacade) DoKTV() {
	fmt.Println("家庭影院 进入KTV 模式")
	hp.tv.On()
	hp.pro.On()
	hp.mp.On()
	hp.Light.Off()
	hp.vb.On()
}

//   游戏 模式 组合
func (hp *HomePlayerFacade) DoGame() {
	fmt.Println("家庭影院 进入 游戏 模式")
	hp.tv.On()
	hp.Light.On()
	hp.xbox.On()
}

func main() {
	homePlayer := new(HomePlayerFacade)

	homePlayer.DoKTV()

	fmt.Println("========================")

	homePlayer.DoGame()
}
