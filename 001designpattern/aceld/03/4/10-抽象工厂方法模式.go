/*
3.3 抽象工厂方法模式
工厂方法模式通过引入工厂等级结构，解决了简单工厂模式中工厂类职责太重的问题，但由于工厂方法模式中的每个工厂只生产一类产品，可能会导致系统中存在大量的工厂类，势必会增加系统的开销。因此，可以考虑将一些相关的产品组成一个“产品族”，由同一个工厂来统一生产，这就是本文将要学习的抽象工厂模式的基本思想。

3.3.1 产品族与产品等级结构
上图表示“产品族”和“产品登记结构”的关系。
产品族：具有同一个地区、同一个厂商、同一个开发包、同一个组织模块等，但是具备不同特点或功能的产品集合，称之为是一个产品族。
产品等级结构：具有相同特点或功能，但是来自不同的地区、不同的厂商、不同的开发包、不同的组织模块等的产品集合，称之为是一个产品等级结构。
当程序中的对象可以被划分为产品族和产品等级结构之后，那么“抽象工厂方法模式”才可以被适用。
“抽象工厂方法模式”是针对“产品族”进行生产产品，具体如下图所示。

3.3.2 抽象工厂模式的角色和职责
抽象工厂（Abstract Factory）角色：它声明了一组用于创建一族产品的方法，每一个方法对应一种产品。
具体工厂（Concrete Factory）角色：它实现了在抽象工厂中声明的创建产品的方法，生成一组具体产品，这些产品构成了一个产品族，每一个产品都位于某个产品等级结构中。
抽象产品（Abstract Product）角色：它为每种产品声明接口，在抽象产品中声明了产品所具有的业务方法。
具体产品（Concrete Product）角色：它定义具体工厂生产的具体产品对象，实现抽象产品接口中声明的业务方法。

可以看出来具体的工厂1，只负责生成具体的产品A1和B1，具体的工厂2，只负责生成具体的产品A2和B2。
“工厂1、A1、B1”为一组，是一个产品族， “工厂2、A2、B2”为一组，也是一个产品族。
3.3.3 抽象工厂方法模式的实现
抽象工厂方法模式按照本章节的案例，可以得到类图如下：

3.3.4 抽象工厂模式的优缺点
优点：
1.  拥有工厂方法模式的优点
2. 当一个产品族中的多个对象被设计成一起工作时，它能够保证客户端始终只使用同一个产品族中的对象。
3   增加新的产品族很方便，无须修改已有系统，符合“开闭原则”。

缺点：
1. 增加新的产品等级结构麻烦，需要对原有系统进行较大的修改，甚至需要修改抽象层代码，这显然会带来较大的不便，违背了“开闭原则”。

3.3.5 适用场景
(1) 系统中有多于一个的产品族。而每次只使用其中某一产品族。可以通过配置文件等方式来使得用户可以动态改变产品族，也可以很方便地增加新的产品族。
(2) 产品等级结构稳定。设计完成之后，不会向系统中增加新的产品等级结构或者删除已有的产品等级结构。

练习：
	设计一个电脑主板架构，电脑包括（显卡，内存，CPU）3个固定的插口，显卡具有显示功能（display，功能实现只要打印出意义即可），内存具有存储功能（storage），cpu具有计算功能（calculate）。
	现有Intel厂商，nvidia厂商，Kingston厂商，均会生产以上三种硬件。
	要求组装两台电脑，
			    1台（Intel的CPU，Intel的显卡，Intel的内存）
			    1台（Intel的CPU， nvidia的显卡，Kingston的内存）
	用抽象工厂模式实现。

package main
// ###### 抽象层
// ###### 基础模块层
// ###### 业务层
func main() {
}

*/

package main

import "fmt"

// 产品等级结构：要求固定。CPU Video MEM

// ###### 抽象层
// ====== 抽象产品
type AbstractCPU interface {
	Calculate()
}
type AbstractVideo interface {
	Display()
}
type AbstractMEM interface {
	Storage()
}

// ====== 抽象工厂
type AbstractFactory interface {
	// 返回抽象的
	CreatCPU() AbstractCPU
	CreatVideo() AbstractVideo
	CreatMEM() AbstractMEM
}

// ###### 基础模块层|实现层
// ====== Intel 产品族
type IntelCPU struct{}

func (i *IntelCPU) Calculate() {
	fmt.Println("英特尔 CPU Calculate")
}

type IntelVideo struct{}

func (i *IntelVideo) Display() {
	fmt.Println("英特尔 显卡 Display")
}

type IntelMEM struct{}

func (i *IntelMEM) Storage() {
	fmt.Println("英特尔 内存 Storage")
}

// ------ Intel 工厂
type IntelFactory struct{}

func (ifa *IntelFactory) CreatCPU() AbstractCPU {
	var cpu AbstractCPU // 抽象的CPU
	cpu = new(IntelCPU)
	return cpu
}

func (ifa *IntelFactory) CreatVideo() AbstractVideo {
	var video AbstractVideo // 抽象的
	video = new(IntelVideo)
	return video
}

func (ifa *IntelFactory) CreatMEM() AbstractMEM {
	var mem AbstractMEM // 抽象的
	mem = new(IntelMEM)
	return mem
}

// ====== Arm 产品族
type ArmCPU struct{}

func (a *ArmCPU) Calculate() {
	fmt.Println("Arm CPU Calculate")
}

type ArmVideo struct{}

func (a *ArmVideo) Display() {
	fmt.Println("Arm 显卡 Display")
}

type ArmMEM struct{}

func (a *ArmMEM) Storage() {
	fmt.Println("Arm 内存 Storage")
}

// ------ Arm 工厂
type ArmFactory struct{}

func (afa *ArmFactory) CreatCPU() AbstractCPU {
	var cpu AbstractCPU // 抽象的CPU
	cpu = new(ArmCPU)
	return cpu
}

func (afa *ArmFactory) CreatVideo() AbstractVideo {
	var video AbstractVideo // 抽象的
	video = new(ArmVideo)
	return video
}

func (afa *ArmFactory) CreatMEM() AbstractMEM {
	var mem AbstractMEM // 抽象的
	mem = new(ArmMEM)
	return mem
}

// ###### 业务层
func main() {
	// A 需要 Intel CPU Video MEM
	// A-1 创建 Intel 工厂
	var iFa AbstractFactory // 虚拟的
	iFa = new(IntelFactory) // 指向具体的工厂

	// A-2 生产 Intel 三样东西
	var iCPU AbstractCPU
	iCPU = iFa.CreatCPU()
	iCPU.Calculate()

	var iVideo AbstractVideo
	iVideo = iFa.CreatVideo()
	iVideo.Display()

	var iMEM AbstractMEM
	iMEM = iFa.CreatMEM()
	iMEM.Storage()

}
