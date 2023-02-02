/*
流程相似，细节 步骤调用不同方法
咖啡： 煮水 冲咖啡 倒入杯中 加糖牛奶
茶： 煮水 冲茶叶 倒入杯中 加柠檬

制作饮料，抽象出来：准备 煮开水 冲泡 倒入杯中 加料

5.1 模板方法模式
5.1.1 模板方法模式中的角色和职责
AbstractClass（抽象类）：在抽象类中定义了一系列基本操作(PrimitiveOperations)，这些基本操作可以是具体的，也可以是抽象的，每一个基本操作对应算法的一个步骤，在其子类中可以重定义或实现这些步骤。同时，在抽象类中实现了一个模板方法(Template Method)，用于定义一个算法的框架，模板方法不仅可以调用在抽象类中实现的基本方法，也可以调用在抽象类的子类中实现的基本方法，还可以调用其他对象中的方法。
ConcreteClass（具体子类）：它是抽象类的子类，用于实现在父类中声明的抽象基本操作以完成子类特定算法的步骤，也可以覆盖在父类中已经实现的具体基本操作。

5.1.3 模板方法的优缺点
优点：
(1) 在父类中形式化地定义一个算法，而由它的子类来实现细节的处理，在子类实现详细的处理算法时并不会改变算法中步骤的执行次序。
(2) 模板方法模式是一种代码复用技术，它在类库设计中尤为重要，它提取了类库中的公共行为，将公共行为放在父类中，而通过其子类来实现不同的行为，它鼓励我们恰当使用继承来实现代码复用。
(3) 可实现一种反向控制结构，通过子类覆盖父类的钩子方法来决定某一特定步骤是否需要执行。
(4) 在模板方法模式中可以通过子类来覆盖父类的基本方法，不同的子类可以提供基本方法的不同实现，更换和增加新的子类很方便，符合单一职责原则和开闭原则。

缺点：
需要为每一个基本方法的不同实现提供一个子类，如果父类中可变的基本方法太多，将会导致类的个数增加，系统更加庞大，设计也更加抽象。

5.1.4 适用场景
(1)具有统一的操作步骤或操作过程;
(2) 具有不同的操作细节;
(3) 存在多个具有同样操作步骤的应用场景，但某些具体的操作细节却各不相同;
在抽象类中统一操作步骤，并规定好接口；让子类实现接口。这样可以把各个具体的子类和操作步骤解耦合。

*/
package main

import "fmt"

type Beverage interface {
	// 煮开水
	BoilWater()
	// 冲泡
	Brew()
	// 倒入杯中
	PourInCup()
	// 加料
	AddThings()
	// 是否加作料的Hook,想加，就重写，返回true。不想加，重写返回false
	WantAddTings() bool
}

// go 接口没有属性。
// 封装一套流程模板基类，让具体制作流程 继承+实现
type template struct {
	b Beverage
}

// 封装的固定模板
func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	// 固定的流程，子类 再去 多态 实现
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()
	if t.b.WantAddTings() {
		t.b.AddThings()
	}

}

// 子类 具体的制作流程
// 咖啡
type MakeCoffee struct {
	template // 继承 模板
}

func (m *MakeCoffee) BoilWater() {
	fmt.Println("水 煮开 100度")
}
func (m *MakeCoffee) Brew() {
	fmt.Println("冲咖啡豆")
}
func (m *MakeCoffee) PourInCup() {
	fmt.Println("咖啡 倒入杯中")
}
func (m *MakeCoffee) AddThings() {
	fmt.Println("杯中 加奶加糖")
}
func (m *MakeCoffee) WantAddTings() bool {
	return true
}

// 茶叶
type MakeTea struct {
	template // 继承 模板
}

func (m *MakeTea) BoilWater() {
	fmt.Println("水 煮 80 度")
}
func (m *MakeTea) Brew() {
	fmt.Println("冲茶叶")
}
func (m *MakeTea) PourInCup() {
	fmt.Println("茶水 倒入杯中")
}
func (m *MakeTea) AddThings() {
	fmt.Println("茶杯中 加 柠檬")
}
func (m *MakeTea) WantAddTings() bool {
	return false
}

// ------ 构造
func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)
	// b 是 Beverage ，是MakeCoffee的接口，这里需要给接口复制。让b 指向子类。来触发b的全部方法的多态特性
	makeCoffee.template.b = makeCoffee
	return makeCoffee
}
func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}

func main() {

	mk := NewMakeCoffee()
	mk.MakeBeverage()
	fmt.Println("---------------------")
	mt := NewMakeTea()
	mt.MakeBeverage()

	// var b Beverage = new(template{MakeCoffee{}})

}
