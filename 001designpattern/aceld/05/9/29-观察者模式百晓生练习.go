package main

import (
	"fmt"
)

/*
               百晓生
	[丐帮]               [明教]
    洪七公               张无忌
    黄蓉					韦一笑
    乔峰				    金毛狮王

	拍手 忽略 复仇
*/

type Event struct {
	Noti    Notifier // 被知晓的通知者
	One     Listener // 事件发起人
	Another Listener // 事件被动的一方
	Msg     string
}

//-------- 抽象层 -------
type Listener interface {
	// 同伴被打
	OnFriendBeFight(event *Event)

	Title() string

	GetName() string
	GetParty() string
}

type Notifier interface {
	Attach(listener Listener) //  AddListener
	Detach(listener Listener) // DeleteListener
	Notify(event *Event)
}

//-------- 实现层 -------
// 观察者 英雄
type Hero struct {
	Name  string
	Party string // 帮派
}

func (h *Hero) OnFriendBeFight(event *Event) {
	// 当事人 忽略
	if h.Name == event.One.GetName() || h.Name == event.Another.GetName() {
		return
	}
	// 打了 别人帮派 拍手
	if h.Party == event.One.GetParty() {
		fmt.Println(h.Title(), "得知消息，拍手较好")
		return
	}
	// 本帮 被打 报仇
	if h.Party == event.Another.GetParty() {
		fmt.Println(h.Title(), "得知消息，要报仇")
		h.Fight(event.One, event.Noti)
		return
	}
}

func (h *Hero) Title() string {
	return fmt.Sprintf("[%s]%s", h.Party, h.Name)
}
func (h *Hero) GetName() string {
	return h.Name
}
func (h *Hero) GetParty() string {
	return h.Party
}

// 主动攻击别人
func (h *Hero) Fight(another Listener, baiXiao Notifier) {
	// 生成武林事件
	event := new(Event)
	event.Msg = fmt.Sprintf("%s 殴打了 %s", h.Title(), another.Title())
	event.Noti = baiXiao
	event.One = h
	event.Another = another

	// 让百晓生知道
	baiXiao.Notify(event)
}

// 被观察者 百晓生
type BaiXiao struct {
	heroList []Listener
}

func (b *BaiXiao) Attach(listener Listener) {
	b.heroList = append(b.heroList, listener)
}
func (b *BaiXiao) Detach(listener Listener) {
	for index, l := range b.heroList {
		if listener == l {
			b.heroList = append(b.heroList[:index], b.heroList[index+1:]...)
		}
	}
}
func (b *BaiXiao) Notify(event *Event) {
	fmt.Println("[世界消息] 百晓生广播、、、", event.Msg)
	for _, listener := range b.heroList {
		listener.OnFriendBeFight(event)
	}
}

// --------业务员逻辑
func main() {
	hero1 := Hero{"黄蓉", "丐帮"}
	hero2 := Hero{"洪七公", "丐帮"}
	hero3 := Hero{"乔峰", "丐帮"}

	hero4 := Hero{"张无忌", "明教"}
	hero5 := Hero{"金毛狮王", "明教"}
	hero6 := Hero{"韦一笑", "明教"}

	baixiao := BaiXiao{}
	baixiao.Attach(&hero1)
	baixiao.Attach(&hero2)
	baixiao.Attach(&hero3)
	baixiao.Attach(&hero4)
	baixiao.Attach(&hero5)
	baixiao.Attach(&hero6)

	fmt.Println("江湖构建完成 ，武林 平静。。。。")

	hero1.Fight(&hero4, &baixiao)
}
