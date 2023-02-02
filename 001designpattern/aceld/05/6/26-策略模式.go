package main

import "fmt"

// ====== 抽象策略 武器策略
type WeaponStrategy interface {
	UseWeapon()
}

// ====== 具体的策略
type Ak47 struct{}

func (a *Ak47) UseWeapon() {
	fmt.Println("用 AK47 扫射")
}

// ====== 具体的策略2
type Knife struct{}

func (k *Knife) UseWeapon() {
	fmt.Println("用 匕首 刺杀")
}

// ====== 环境类 使用策略的。当接口用
type Hero struct {
	strategy WeaponStrategy // 拥有一个抽象的策略、
}

//  ====== 赋值，设计一个策略的方法
func (h *Hero) SetWeaponStrategy(s WeaponStrategy) {
	h.strategy = s
}

// ====== 业务 战斗方法
func (h *Hero) Fight() {
	h.strategy.UseWeapon() // 调用具体的策略
}

func main() {
	hero := Hero{}

	// 策略1
	hero.SetWeaponStrategy(new(Ak47))
	hero.Fight()

	// 策略1
	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()
}
