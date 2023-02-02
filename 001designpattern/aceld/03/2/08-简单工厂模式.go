/*
3.1.2 简单工厂模式角色和职责
	简单工厂模式并不属于GoF的23种设计模式。他是开发者自发认为的一种非常简易的设计模式，其角色和职责如下：
	工厂（Factory）角色：简单工厂模式的核心，它负责实现创建所有实例的内部逻辑。工厂类可以被外界直接调用，创建所需的产品对象。
	抽象产品（AbstractProduct）角色：简单工厂模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。
	具体产品（Concrete Product）角色：简单工厂模式所创建的具体实例对象。

      上述代码可以看出，业务逻辑层只会和工厂模块进行依赖，这样业务逻辑层将不再关心Fruit类是具体怎么创建基础对象的。

	3.1.4 简单工厂方法模式的优缺点
	优点：
	1. 实现了对象创建和使用的分离。
	2. 不需要记住具体类名，记住参数即可，减少使用者记忆量。

	缺点：
	1. 对工厂类职责过重，一旦不能工作，系统受到影响。
	2. 增加系统中类的个数，复杂度和理解度增加。
	3. 违反“开闭原则”，添加新产品需要修改工厂逻辑，工厂越来越复杂。

	适用场景：
	1.  工厂类负责创建的对象比较少，由于创建的对象较少，不会造成工厂方法中的业务逻辑太过复杂。
	2. 客户端只知道传入工厂类的参数，对于如何创建对象并不关心。
*/

package main

import "fmt"

// ###### 抽象层
type Fruit interface {
	Show() // 接口的方法
}

// ###### 实现层
type Apple struct {
	Fruit // 为了便于理解
}

func (a *Apple) Show() {
	fmt.Println("苹果 show 了")

}

type Banana struct {
	Fruit // 为了便于理解
}

func (b *Banana) Show() {
	fmt.Println("香蕉 show 了")
}

// ++++++ 工厂模块
type Factory struct {
}

// 返回抽象的Fruit
func (f *Factory) CreatFruit(kind string) Fruit {

	var fruit Fruit

	if kind == "apple" {
		// apple 的初始化业务
		fruit = new(Apple) // 满足多态条件赋值，父类指针，指向子类对象
	} else if kind == "banana" {
		fruit = new(Banana) // 满足多态条件赋值，父类指针，指向子类对象
	}

	return fruit

}

// ###### 业务层
func main() {
	// 这里操作的，都是 Fruit 类型的
	// main 一直面向抽象层 开发
	factory := new(Factory)
	var apple Fruit
	apple = factory.CreatFruit("apple")
	apple.Show()
	banana := factory.CreatFruit("banana")
	banana.Show()
}
