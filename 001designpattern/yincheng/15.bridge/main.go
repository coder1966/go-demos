package main

import "godemos/001designpattern/yincheng/15.bridge/bridge"

// SMS
// Email
func main() {
	// m1 := bridge.NewCommonMessage(bridge.ViaSMS())
	// m2 := bridge.NewCommonMessage(bridge.ViaEmail())
	m1 := bridge.NewUrgencyMessage(bridge.ViaSMS())
	m2 := bridge.NewUrgencyMessage(bridge.ViaEmail())
	m1.SendMessage("baby 你好", "babymm")
	m2.SendMessage("hi 哥们", "vencent")

	m3 := bridge.NewUrgencyMessage(bridge.ViaEmail())
	m3.SendMessage("good 老大", "leader")
}
