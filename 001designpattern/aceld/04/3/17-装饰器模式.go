/*
4.2 装饰模式
装饰模式(Decorator Pattern)：动态地给一个对象增加一些额外的职责，就增加对象功能来说，装饰模式比生成子类实现更为灵活。装饰模式是一种对象结构型模式。

4.2.1 装饰模式中的角色和职责
Component（抽象构件）：它是具体构件和抽象装饰类的共同父类，声明了在具体构件中实现的业务方法，它的引入可以使客户端以一致的方式处理未被装饰的对象以及装饰之后的对象，实现客户端的透明操作。
ConcreteComponent（具体构件）：它是抽象构件类的子类，用于定义具体的构件对象，实现了在抽象构件中声明的方法，装饰器可以给它增加额外的职责（方法）。
其标准的类图如下所示：

4.2.3 装饰模式的优缺点
优点：
(1) 对于扩展一个对象的功能，装饰模式比继承更加灵活性，不会导致类的个数急剧增加。
(2) 可以通过一种动态的方式来扩展一个对象的功能，从而实现不同的行为。
(3) 可以对一个对象进行多次装饰。
(4) 具体构件类与具体装饰类可以独立变化，用户可以根据需要增加新的具体构件类和具体装饰类，原有类库代码无须改变，符合“开闭原则”。
缺点：
(1) 使用装饰模式进行系统设计时将产生很多小对象，大量小对象的产生势必会占用更多的系统资源，影响程序的性能。
(2) 装饰模式提供了一种比继承更加灵活机动的解决方案，但同时也意味着比继承更加易于出错，排错也很困难，对于多次装饰的对象，调试时寻找错误可能需要逐级排查，较为繁琐。

4.2.3 适用场景
(1) 动态、透明的方式给单个对象添加职责。
(2) 当不能采用继承的方式对系统进行扩展或者采用继承不利于系统扩展和维护时可以使用装饰模式。
装饰器模式关注于在一个对象上动态的添加方法，然而代理模式关注于控制对对象的访问。换句话说，用代理模式，代理类（proxy class）可以对它的客户隐藏一个对象的具体信息。因此，当使用代理模式的时候，我们常常在一个代理类中创建一个对象的实例。并且，当我们使用装饰器模式的时候，我们通常的做法是将原始对象作为一个参数传给装饰者的构造器。
*/
package main

import "fmt"

// ====== 抽象层 Component（抽象构件）==========
type Phone interface {
	Show() // 构建的功能
}

// ++++++抽象的装饰器:装饰器基础类，本来应该interface，但是go interface不支持属性
type Decorator struct {
	phone Phone // 组合进来抽象的phone
}

// ++++++抽象的方法, 空的，其他具体装饰器炫耀重写这个方法
func (d *Decorator) Show() {}

// ====== 实现层 ConcreteComponent（具体构件==========
// ++++++具体构建
type Huawei struct{}

func (hw *Huawei) Show() {
	fmt.Println("show huawei 手机")
}

type Xiaomi struct{}

func (xm *Xiaomi) Show() {
	fmt.Println("show Xiaomi 手机")
}

// ++++++具体装饰器
// ------贴膜的
type MoDecorator struct {
	Decorator // 继承 基础装饰器累（主要为了phone的成员属性）
}

func (md *MoDecorator) Show() {
	md.phone.Show()    // 先调原来的被装饰构建的方法
	fmt.Println("贴膜了") // 额外的装饰功能
}

// 构造函数
func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone: phone}}
}

// ------手机壳
type KeDecorator struct {
	Decorator // 继承 基础装饰器累（主要为了phone的成员属性）
}

func (kd *KeDecorator) Show() {
	kd.phone.Show()     // 先调原来的被装饰构建的方法
	fmt.Println("加外壳了") // 额外的装饰功能
}

// 构造函数
func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone: phone}}
}

// ====== 业务逻辑层 ==========
func main() {
	var huawei Phone = new(Huawei)
	huawei.Show()

	// 通过 膜 的 Decorator 对 对象 装饰
	var moHuawei Phone = NewMoDecorator(huawei)
	moHuawei.Show()

	var keHuawei Phone = NewKeDecorator(huawei)
	keHuawei.Show()

	// 多级装饰
	// 通过 膜 的 Decorator 对 已经 壳 装饰过的 对象 装饰
	var keMoHuawei Phone = NewMoDecorator(keHuawei)
	keMoHuawei.Show()
}
