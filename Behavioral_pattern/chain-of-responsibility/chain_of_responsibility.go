package chain_of_responsibility

import "fmt"

type Component interface {
	SetNext(component Component) Component
	Handle(request string)
}

type BaseComponent struct {
	NextComponent Component
}

func (c *BaseComponent) SetNext(component Component) Component {
	c.NextComponent = component
	return component
}

func (c *BaseComponent) Handle(request string) {
	if c.NextComponent != nil {
		c.NextComponent.Handle(request)
		return
	}
	fmt.Println("BaseComponent handle")
}

func NewBaseComponent() *BaseComponent {
	return &BaseComponent{}
}

type ConcreteComponentA struct {
	BaseComponent
}

func (c *ConcreteComponentA) Handle(request string) {
	if request == "ConcreteComponentA" {
		fmt.Println("ConcreteComponentA handle")
	} else {
		c.NextComponent.Handle(request)
	}
}

func NewConcreteComponentA() *ConcreteComponentA {
	return &ConcreteComponentA{}
}

type ConcreteComponentB struct {
	BaseComponent
}

func (c *ConcreteComponentB) Handle(request string) {
	if request == "ConcreteComponentB" {
		fmt.Println("ConcreteComponentB handle")
	} else {
		c.NextComponent.Handle(request)
	}
}

func NewConcreteComponentB() *ConcreteComponentB {
	return &ConcreteComponentB{}
}
