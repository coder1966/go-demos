package main

import "fmt"

// ------ 命令接受者 厨师
type Cooker struct{}

func (c Cooker) MakeChicken() {
	fmt.Println("厨师  烤了 鸡翅")
}
func (c Cooker) MakeSheep() {
	fmt.Println("厨师  烤了 羊肉串")
}

// 抽象的命令
type Command interface {
	Treat() // 抽象的指令
}

// 菜单，鸡翅
type CommandMakeChicken struct {
	cooker *Cooker
}

func (c *CommandMakeChicken) Treat() {
	c.cooker.MakeChicken()
}

// 菜单，羊肉
type CommandMakeSheep struct {
	cooker *Cooker
}

func (c *CommandMakeSheep) Treat() {
	c.cooker.MakeSheep()
}

// ===== 命令调用者 服务生
type Waiter struct {
	cmdList []Command // 抽象命令
}

// ====== 发送命令
func (w *Waiter) Notify() {
	for _, cmd := range w.cmdList {
		cmd.Treat() // 多态 执行
	}
}

// 吃客
func main() {
	// 厨师
	cooker := new(Cooker)
	// 菜单
	cmdChicken := CommandMakeChicken{cooker}
	cmdSheep := CommandMakeSheep{cooker}
	// 服务生
	waiter := new(Waiter)
	waiter.cmdList = append(waiter.cmdList, &cmdChicken, &cmdSheep)
	waiter.Notify()

}
