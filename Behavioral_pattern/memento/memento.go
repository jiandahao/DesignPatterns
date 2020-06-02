package memento

import (
	"fmt"
	"time"
)

//主要思想：在不暴露对象实现细节的情况下保存和恢复对象之前的状态

// Don't expose memento class
// 定义一个备忘录类，该类不能在包（package memento）外直接使用
// 其成员变量也都是私有的。但可以定义一些对外暴露的接口，例如该例中
// 的GetMementoName和GetMementoDate
type memento struct {
	snapName string
	snapDate time.Time
	state    string
}

func (m *memento) GetMementoName() string {
	return m.snapName
}

func (m *memento) GetMementoDate() time.Time {
	return m.snapDate
}

type Originator struct {
	name string // 不暴露该字段，隐藏具体实现细节
}

// 唯一能够使用memento对象的方式是，通过CreateMemento获取
func (o *Originator) CreateMemento(name string) *memento {
	fmt.Println("make memento")
	return &memento{
		snapDate: time.Now(),
		snapName: name,
		state:    o.name,
	}
}

func (o *Originator) Restore(snapshot *memento) {
	fmt.Println("Restore")
	o.name = snapshot.state
}

func (o *Originator) SetName(name string) {
	o.name = name
}

// 定义一个负责人，负责人仅知道 “何时” 和 “为何” 捕捉原发器的状态， 以及何时恢复状态。
// 负责人通过保存备忘录栈来记录原发器的历史状态。 当原发器需要回溯历史状态时，
// 负责人将从栈中获取最顶部的备忘录， 并将其传递给原发器的恢复 （restoration） 方法。
type Caretaker struct {
	originator *Originator
	history    []*memento
}

func (taker *Caretaker) ChangeName(name string) {
	if taker.originator == nil {
		taker.originator = &Originator{}
	}
	snap := taker.originator.CreateMemento(name)
	taker.originator.SetName(name)
	taker.history = append(taker.history, snap)
}

func (taker *Caretaker) Undo() {
	if taker == nil || len(taker.history) <= 0 {
		return
	}
	snap := taker.history[len(taker.history)-1]
	taker.history = taker.history[:len(taker.history)-1]
	taker.originator.Restore(snap)
}

func NewCaretaker(originator *Originator) *Caretaker {
	return &Caretaker{
		originator: originator,
		history:    []*memento{},
	}
}
