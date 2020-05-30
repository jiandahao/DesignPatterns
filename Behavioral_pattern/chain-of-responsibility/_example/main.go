package main

import cor "github.com/jiandahao/DesignPatterns/Behavioral_pattern/chain-of-responsibility"

func main() {
	baseComponent := cor.NewBaseComponent()
	componentA := cor.NewConcreteComponentA()
	componentB := cor.NewConcreteComponentB()
	componentA.SetNext(componentB).SetNext(baseComponent)

	componentA.Handle("ConcreteComponentA")
	componentA.Handle("ConcreteComponentB")
	componentA.Handle("Base")
}
