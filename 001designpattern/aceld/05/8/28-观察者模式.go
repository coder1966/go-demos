package main

import "fmt"

//--------- 抽象层 --------
// 抽象的观察者
type Listener interface {
	OnTeacherComing() // 观察者 得到 触发 的具体动作
}

// 抽象的通知者 被观察目标
type Notifier interface {
	Attach(listener Listener) //  AddListener
	Detach(listener Listener) // DeleteListener
	Notify()                  // 通知全部
}

//--------- 实现层 --------
// 观察者 具体的学生
type StuZhangSan struct {
	BadThing string
}

func (s StuZhangSan) OnTeacherComing() {
	fmt.Println("张三 停止 了 ", s.BadThing)
}

// 不是多态，不需要抽象
func (s StuZhangSan) DoBadthing() {
	fmt.Println("张三 正在 ", s.BadThing)
}

type StuLiSi struct {
	BadThing string
}

func (s StuLiSi) OnTeacherComing() {
	fmt.Println("李四 加剧 了 ", s.BadThing)
}

// 不是多态，不需要抽象
func (s StuLiSi) DoBadthing() {
	fmt.Println("李四 正在 ", s.BadThing)
}

type StuWangWu struct {
	BadThing string
}

func (s StuWangWu) OnTeacherComing() {
	fmt.Println("王五 停止 了 ", s.BadThing)
}

// 不是多态，不需要抽象
func (s StuWangWu) DoBadthing() {
	fmt.Println("王五 正在 ", s.BadThing)
}

// Subject 通知者 班长
type ClassMonitor struct {
	listenerList []Listener // 抽象的 ，被通知的
}

// 加
func (c *ClassMonitor) Attach(listener Listener) {
	c.listenerList = append(c.listenerList, listener)
}

// 减
func (c *ClassMonitor) Detach(listener Listener) {
	// c.listenerList = append(c.listenerList, listener)
	for i, l := range c.listenerList {
		if listener == l {
			// aa := c.listenerList[:i]
			// aa = c.listenerList[i+1:]
			// _ = aa
			c.listenerList = append(c.listenerList[:i], c.listenerList[i+1:]...)
			break
		}
	}
}

// 通知
func (c *ClassMonitor) Notify() {
	for _, l := range c.listenerList {
		l.OnTeacherComing() // 被通知学生，动作
	}
}

//--------- 实现 --------
func main() {
	s1 := &StuZhangSan{
		BadThing: "抄作业",
	}
	s2 := &StuLiSi{
		BadThing: "唱歌",
	}
	s3 := &StuWangWu{
		BadThing: "化妆",
	}

	// 班长
	// cm := &ClassMonitor{}
	cm := new(ClassMonitor)
	// 加观察者
	cm.Attach(s1)
	cm.Attach(s2)
	cm.Attach(s3)

	// 学生嗨皮
	s1.DoBadthing()
	s2.DoBadthing()
	s3.DoBadthing()

	fmt.Println("老师来了，班长通知")
	cm.Notify()

	fmt.Println("-----删除 lisi")
	cm.Detach(s2)
	cm.Detach(s2) // 重复，看报错不
	cm.Notify()

}
