package main

import "fmt"

// 弄一个责任链，代替一大堆if

// https://www.bilibili.com/video/BV13A411U7Z4/?p=16&spm_id_from=pageDriver

type SellInfo struct {
	Price      float64
	OrderCount int
	TotalCount int
	MemberShip int
}

type Rule interface {
	Check(sellInfo *SellInfo) bool
}

func Chain(sellInfo *SellInfo, rules ...Rule) bool {
	for _, r := range rules {
		if !r.Check(sellInfo) {
			return false
		}
	}
	return true
}

func main() {
	a := &SellInfo{
		Price:      1.9,
		OrderCount: 1,
		TotalCount: 20,
		MemberShip: 1,
	}

	rules := []Rule{
		&PriceRule{},
		&OrderCountRule{},
		&MemberShipRule{},
		&DiscountRule{},
		// ...
	}

	r := Chain(a, rules...)
	fmt.Println("结果：", r)
}

type PriceRule struct{}

func (pr *PriceRule) Check(sellInfo *SellInfo) bool {
	return sellInfo.Price > 0
}

type OrderCountRule struct{}

func (pr *OrderCountRule) Check(sellInfo *SellInfo) bool {
	return sellInfo.TotalCount > sellInfo.OrderCount
}

type MemberShipRule struct{}

func (pr *MemberShipRule) Check(sellInfo *SellInfo) bool {
	return sellInfo.MemberShip == 1
}

type DiscountRule struct{}

func (pr *DiscountRule) Check(sellInfo *SellInfo) bool {
	return sellInfo.Price < 100 && sellInfo.MemberShip == 2
}
