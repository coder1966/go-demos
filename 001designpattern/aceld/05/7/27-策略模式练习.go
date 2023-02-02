package main

import "fmt"

/*
	练习：
	商场促销有策略A（0.8折）策略B（消费满200，返现100），用策略模式模拟场景
*/

// ====== 抽象策略 销售策略
type SellStrategy interface {
	GetPrice(price float64) float64 // 根据原价，打折 到 售卖价
}

// ====== 具体的策略01
type StrategyA08 struct{}

func (s *StrategyA08) GetPrice(price float64) float64 {
	fmt.Println("策略 A 打八折")
	return price * 0.8
}

// ====== 具体的策略2
type StrategyB200_100 struct{}

func (s *StrategyB200_100) GetPrice(price float64) float64 {
	fmt.Println("策略 B 满200 减 100")
	if price >= 200 {
		return price - 100
	}
	return price
}

// ====== 环境类 使用策略 当接口用
type Goods struct {
	Price    float64
	Strategy SellStrategy // 拥有有个抽象策略
}

// 赋值 更换 设计一个策略方法
func (g *Goods) SetSellStrategy(s SellStrategy) {
	g.Strategy = s
}

// 业务 打折方法
func (g *Goods) SellPrice() float64 {
	fmt.Println("原价格： ", g.Price)
	return g.Strategy.GetPrice(g.Price)
}

func main() {
	nike := Goods{
		Price: 200,
	}

	// 上午 方法一
	nike.SetSellStrategy(new(StrategyA08))
	fmt.Println("上午 方法一，卖 ", nike.SellPrice())
	// 下午 方法二
	nike.SetSellStrategy(new(StrategyB200_100))
	fmt.Println("下午 方法二，卖 ", nike.SellPrice())
}
