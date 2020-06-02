package mediator

import "fmt"

type Mediator interface {
	Notify(sender Component, event string)
}

type ConcreteMediator struct {
	Button *Button
	CheckBox *CheckBox
}

func (m *ConcreteMediator) Notify(sender Component, event string) {
	switch event {
	case "button":
		fmt.Println("Mediator reacts on button and triggers check box to check")
		m.CheckBox.Check()
	case "check box":
		fmt.Println("Mediator reacts on check box and triggers disable button")
		m.Button.Disable()
	}
}

type Component interface {
	OnClick()
}

type Button struct {
	Component
	mediator Mediator
}

func (b *Button) OnClick() {
	b.mediator.Notify(b, "button")
}

func (b *Button) Disable() {
	fmt.Println("disable button")
}

func NewButton(mediator Mediator) *Button {
	return &Button{mediator:mediator}
}

type CheckBox struct {
	Component
	mediator Mediator
}

func (b *CheckBox) OnClick() {
	b.mediator.Notify(b, "check box")
}

func (b *CheckBox) Check() {
	fmt.Println("Check Box does check")
}

func NewCheckBox(mediator Mediator) *CheckBox {
	return &CheckBox{mediator:mediator}
}