package main

import (
	"fmt"
	"github.com/jiandahao/DesignPatterns/Structural_patterns/adapter"
	"reflect"
)

func ClientCode(target adapter.Target) {
	fmt.Println(reflect.TypeOf(target.Request()))
}
func main() {
	normalTarget := &adapter.NormalTarget{}
	// Client Code works fine with NormalTarget
	ClientCode(normalTarget)

	specialTarget := adapter.Adaptee{}
	fmt.Println(reflect.TypeOf(specialTarget.SpecialRequest()))
	// specialTarget could not be used directly by client
	// ClientCode(specialTarget)

	// using adapter
	specialTargetAdapter := &adapter.Adapter{Adaptee: specialTarget}
	ClientCode(specialTargetAdapter)
}
