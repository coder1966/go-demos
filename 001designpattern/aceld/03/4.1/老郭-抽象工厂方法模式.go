package main

import "fmt"

// 形状接口
type Shape interface {
	Draw()
}

// 色彩接口
type Color interface {
	Fill()
}

// 实现模型接口的圆形类
type Circle struct{}

// 正方形
type Square struct{}

// 圆形 类 的 Draw 方法
func (c *Circle) Draw() {
	fmt.Println("画圆")
}

// 正方形 类 的 Draw 方法
func (s *Square) Draw() {
	fmt.Println("画正方形")
}

// 实现 色彩接口 的红色 类
type Red struct{}

// 实现色彩接口的 绿色 类
type Green struct{}

// 红的类 Fill 方法
func (r *Red) Fill() {
	fmt.Println("填充红色")
}

// 绿色类 Fill
func (g *Green) Fill() {
	fmt.Println("填充绿色")
}

// AbstractFactory 抽象工厂接口
type AbstractFactory interface {
	GetShape(shapeName string) Shape
	GetColor(colorName string) Color
}

// 模型工厂的类
type ShapeFactory struct{}

// 色彩工厂的类
type ColorFactory struct{}

// 模型工厂 实例获得模型子类的方法
func (sf *ShapeFactory) GetShape(shapeName string) Shape {
	switch shapeName {
	case "circle":
		return &Circle{}
	case "square":
		return &Square{}
	default:
		return nil
	}
}

// 模型工厂实例，不需要色彩
func (sf *ShapeFactory) GetColor(shapeName string) Color {
	return nil
}

// 色彩工厂 实例获得具体色彩子类
func (sf *ColorFactory) GetColor(colorName string) Color {
	switch colorName {
	case "red":
		return &Red{}
	case "green":
		return &Green{}
	default:
		return nil
	}
}

// 色彩工厂实例，不需要形状
func (sf *ColorFactory) GetShape(colorName string) Shape {
	return nil
}

// ===========
// 超级工厂类，用来获取工厂实例
type FactoryProducer struct{}

// 获取工厂的方法
func (fp *FactoryProducer) GetFactory(factoryName string) AbstractFactory {
	switch factoryName {
	case "color":
		return &ColorFactory{}
	case "shape":
		return &ShapeFactory{}
	default:
		return nil
	}
}

// 构造函数
func NewFactoryProducer() *FactoryProducer {
	return &FactoryProducer{}
}

// 主程序
func main() {
	superF := NewFactoryProducer()
	colorF := superF.GetFactory("color")
	shapeF := superF.GetFactory("shape")

	red := colorF.GetColor("red")
	green := colorF.GetColor("green")

	circle := shapeF.GetShape("circle")
	square := shapeF.GetShape("square")

	// 红色圆形
	circle.Draw()
	red.Fill()

	// 绿色方形
	square.Draw()
	green.Fill()
}
