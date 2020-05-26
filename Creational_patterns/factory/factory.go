package factory

import "fmt"

// The product interface declares the operations that all
// concrete products must implement.
type Button interface {
	Render()
	OnClick()
}

// Concrete products provide various implementations of the
// product interface.
type WindowButton struct{}

func (b *WindowButton) Render() {
	fmt.Println("Rendering window button")
}

func (b *WindowButton) OnClick() {
	fmt.Println("Clicking window button")
}

type LinuxButton struct{}

func (b *LinuxButton) Render() {
	fmt.Println("Rendering linux button")
}

func (b *LinuxButton) OnClick() {
	fmt.Println("Clicking linux button")
}

// The creator class declares the factory method that must
// return an object of a product class. The creator's subclasses
// usually provide the implementation of this method.
type DialogFactory interface {
	// The creator may also provide some default implementation
	// of the factory method.
	CreateButton() Button
}

type WindowDialogFactory struct{}

// Concrete creators override the factory method to change the
// resulting product's type.
func (b *WindowDialogFactory) CreateButton() Button {
	return &WindowButton{}
}

type LinuxDialogFactory struct{}

func (b *LinuxDialogFactory) CreateButton() Button {
	return &LinuxButton{}
}
