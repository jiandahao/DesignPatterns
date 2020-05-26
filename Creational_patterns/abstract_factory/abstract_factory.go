package abstract_factory

import "fmt"

// 抽象工厂接口声明了一组能返回不同抽象产品的方法。这些产品属于同一个系列
// 且在高层主题或概念上具有相关性。同系列的产品通常能相互搭配使用。系列产
// 品可有多个变体，但不同变体的产品不能搭配使用。
type GUIFactory interface {
	CreateButton() Button
	CreateCheckBox() CheckBox
}

type WindowsGUIFactory struct{}

func (f *WindowsGUIFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (f *WindowsGUIFactory) CreateCheckBox() CheckBox {
	return &WindowsCheckBox{}
}

type LinuxGUIFactory struct{}

func (f *LinuxGUIFactory) CreateButton() Button {
	return &LinuxButton{}
}

func (f *LinuxGUIFactory) CreateCheckBox() CheckBox {
	return &LinuxCheckBox{}
}

type Button interface {
	OnClick()
}

type WindowsButton struct{}
type LinuxButton struct{}

func (b *WindowsButton) OnClick() {
	fmt.Println("click windows button")
}

func (b *LinuxButton) OnClick() {
	fmt.Println("click linux button")
}

type CheckBox interface {
	Check()
}

type WindowsCheckBox struct{}
type LinuxCheckBox struct{}

func (c *WindowsCheckBox) Check() {
	fmt.Println("check windows check box")
}

func (c *LinuxCheckBox) Check() {
	fmt.Println("check linux check box")
}
