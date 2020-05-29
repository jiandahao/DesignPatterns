package main

import (
	"fmt"
	"github.com/jiandahao/DesignPatterns/Structural_patterns/composite"
)

func main() {
	comp := composite.NewComposite()
	dotA := composite.NewDot(1.2, 2.4)
	dotB := composite.NewDot(2.4, 1.2)

	//circle := &composite.Circle{
	//	Radius: 12,
	//	Dot:    dot,
	//}
	fmt.Println("=======Add two dots========")
	comp.AddComponent(dotA)
	comp.AddComponent(dotB)

	fmt.Println("========Draw dots==========")
	comp.Draw()

	// remove dot B
	fmt.Println("=======Remove dot B========")
	comp.RemoveComponent(dotB)
	comp.Draw()

	// could not add self
	comp.AddComponent(comp)

	fmt.Println("=========Move dots=========")
	comp.AddComponent(dotB)
	comp.Move(10, 12)

	// add a composite
	newComp := composite.NewComposite()
	newComp.AddComponent(comp)
	newComp.Draw()
}
