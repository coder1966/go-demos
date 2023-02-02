/*
4.1 代理模式
	Proxy模式又叫做代理模式，是构造型的设计模式之一，它可以为其他对象提供一种代理（Proxy）以控制对这个对象的访问。
	所谓代理，是指具有与代理元（被代理的对象）具有相同的接口的类，客户端必须通过代理与被代理的目标类交互，而代理一般在交互的过程中（交互前后），进行某些特别的处理。
用一个日常可见的案例来理解“代理”的概念，如下图：

这里假设有一个“自己”的角色，正在玩一款网络游戏。称这个网络游戏就是代理模式的“Subject”，表示要做一件事的目标或者对象事件主题。
（1）“自己”有一个给游戏角色升级的需求或者任务，当然“自己”可以独自完成游戏任务的升级。
（2）或者“自己”也可以邀请以为更加擅长游戏的“游戏代练”来完成升级这件事，这个代练就是“Proxy”代理。
（3）“游戏代练”不仅能够完成升级的任务需求，还可以额外做一些附加的能力。比如打到一些好的游戏装备、加入公会等等周边收益。
所以代理的出现实则是为了能够覆盖“自己”的原本的需求，且可以额外做其他功能，这种额外创建的类是不影响已有的“自己”和“网络游戏”的的关系。是额外添加，在设计模式原则上，是符合“开闭原则”思想。那么当需要给“自己”增加额外功能的时候，又不想改变自己，那么就选择邀请一位”代理”来完成吧。

subject（抽象主题角色）：真实主题与代理主题的共同接口。
	RealSubject（真实主题角色）：定义了代理角色所代表的真实对象。
	Proxy（代理主题角色）：含有对真实主题角色的引用，代理角色通常在将客户端调用传递给真是主题对象之前或者之后执行某些操作，而不是单纯返回真实的对象。
4.1.2 代理模式案例实现
讲述标准类图改成一个案例来理解。

这里以一个购物作为一个主题任务，这是一个抽象的任务。具体的购物主题分别包括“韩国购物”、“美国购物”、“非洲购物”等。可以这些都是“自己”去完成主题，那么如果希望不仅完成购物，还要做真假辨别、海关安检等，同样依然能够完成自己本身的具体购物主题，那么则可以创建一个新的代理来完成这件事。代理需要将被代理的主题关联到本类中，去重新实现Buy()方法，在Buy()方法中，调用被调用的Buy()，在额外完成辨别真伪和海关安检两个任务动作，具体的代码实现如下：

4.1.4 代理模式的优缺点
优点：
(1) 能够协调调用者和被调用者，在一定程度上降低了系统的耦合度。
(2) 客户端可以针对抽象主题角色进行编程，增加和更换代理类无须修改源代码，符合开闭原则，系统具有较好的灵活性和可扩展性。

缺点：
(1) 代理实现较为复杂。

4.1.5 适用场景
	为其他对象提供一种代理以控制对这个对象的访问。

package main
// ###### 抽象层
// ###### 基础模块层
// ###### 业务层
func main() {
}

*/

package main

import (
	"fmt"
)

type Goods struct {
	Kind string // 商品种类
	Fact bool   // 真假
}

// ###### 抽象层
type Shopping interface {
	Buy(goos *Goods) // 某任务
}

// ###### 实现层
type KoreaShopping struct{}

func (k *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("韩国 购物 买了", goods.Kind)
}

type AmericanShopping struct{}

func (a *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("美国 购物 买了", goods.Kind)
}

// ======= 海外代理 抽象的
type OverSeaProxy struct {
	shopping Shopping // 代理某个主题，抽象的类型
}

func NewProxy(shopping Shopping) Shopping {
	return &OverSeaProxy{shopping: shopping}
}

func (o *OverSeaProxy) Buy(goods *Goods) {
	// 1 辨别真伪
	if o.distinguish(goods) == true {
		// 2 调用具体购物方式 Buy() 方法
		o.shopping.Buy(goods)
		// 3 海关安检
		o.check(goods)
	}

}

func (o *OverSeaProxy) distinguish(goods *Goods) bool {
	fmt.Println("对 [", goods.Kind, "] 进行了辨别真伪。")
	if goods.Fact == false {
		fmt.Println("假货  ", goods.Kind)
	}
	return goods.Fact
}

func (o *OverSeaProxy) check(goods *Goods) {
	fmt.Println("对 [", goods.Kind, "] 进行了海关安检。成功带回")
}

// ###### 业务层
func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	g2 := Goods{
		Kind: "美国证书",
		Fact: false,
	}

	var kShopping Shopping
	kShopping = new(KoreaShopping)

	var k_proxy Shopping
	k_proxy = NewProxy(kShopping)
	k_proxy.Buy(&g1)

	var aShopping Shopping
	aShopping = new(AmericanShopping)

	var a_proxy Shopping
	a_proxy = NewProxy(aShopping)
	a_proxy.Buy(&g2)

}
