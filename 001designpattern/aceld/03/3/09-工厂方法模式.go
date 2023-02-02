/*
3.2 工厂方法模式
3.2.1 工厂方法模式中的角色和职责
抽象工厂（Abstract Factory）角色：工厂方法模式的核心，任何工厂类都必须实现这个接口。
工厂（Concrete Factory）角色：具体工厂类是抽象工厂的一个实现，负责实例化产品对象。
抽象产品（Abstract Product）角色：工厂方法模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。
具体产品（Concrete Product）角色：工厂方法模式所创建的具体实例对象。

可以看见，新增的基本类“日本苹果”，和“具体的工厂” 均没有改动之前的任何代码。完全符合开闭原则思想。新增的功能不会影响到之前的已有的系统稳定性。

3.2.3 工厂方法模式的优缺点
优点：
1. 不需要记住具体类名，甚至连具体参数都不用记忆。
2. 实现了对象创建和使用的分离。
3. 系统的可扩展性也就变得非常好，无需修改接口和原类。
4.对于新产品的创建，符合开闭原则。

缺点：
1. 增加系统中类的个数，复杂度和理解度增加。
2. 增加了系统的抽象性和理解难度。

适用场景：
1. 客户端不知道它所需要的对象的类。
2. 抽象工厂类通过其子类来指定创建哪个对象。

package main
// ###### 抽象层
// ###### 基础模块层
// ###### 业务层
func main() {
}

*/

package main

import "fmt"

// ###### 抽象层
// ++++++ 水果类（抽象的接口）
type Fruit interface {
	Show()
}

// ++++++ 工厂类（抽象的接口）
type AbstractFactory interface {
	// 水果产生器，返回抽象的指针
	CreatFruit() Fruit
}

// ###### 基础模块层
// ++++++ 水果模块
// ------ 苹果
type Apple struct {
	Fruit // 为了便于理解
}

func (a *Apple) Show() {
	fmt.Println("苹果 show 了")

}

// ------ 香蕉
type Banana struct {
	Fruit // 为了便于理解
}

func (b *Banana) Show() {
	fmt.Println("香蕉 show 了")
}

// ------ 可增加 水果
// ------ 可增加 水果
// ------ 可增加 水果

// ++++++ 工厂模块
// ------ 具体的苹果工厂
type AppleFactory struct {
	AbstractFactory // 可以省略，这里显式了
}

func (a *AppleFactory) CreatFruit() Fruit {
	var fruit Fruit // 虚拟的指针

	// 生产具体的苹果
	fruit = new(Apple)
	return fruit
}

// ------ 具体的香蕉工厂
type BananaFactory struct {
	AbstractFactory // 可以省略，这里显式了
}

func (b *BananaFactory) CreatFruit() Fruit {
	var fruit Fruit // 虚拟的指针

	// 生产具体的香蕉
	fruit = new(Banana)
	return fruit
}

// ------ 可增加 工厂
// ------ 可增加 工厂
// ------ 可增加 工厂

// ###### 业务层
func main() {
	// A：具体的苹果对象
	// A-1先来一个具体的苹果工厂
	var appleFac AbstractFactory
	appleFac = new(AppleFactory)
	// A-2生产一个具体的水果
	var apple Fruit
	apple = appleFac.CreatFruit()

	apple.Show() // 发生了 多态。apple 是 虚拟的 水果类
}
