package main

import (
	"fmt"
	"github.com/jiandahao/DesignPatterns/Creational_patterns/prototype"
)

func main() {
	var protoTypeList []prototype.ProtoType

	protoTypeA := &prototype.ConcreteProtoTypeA{
		Name: "test name",
		Age:  10,
	}

	protoTypeList = append(protoTypeList, protoTypeA)
	protoTypeList = append(protoTypeList, protoTypeA.Clone())

	protoTypeB := &prototype.ConcreteProtoTypeB{
		Nickname:    "test nickname",
		PhoneNumber: "+86123456789",
	}

	protoTypeList = append(protoTypeList, protoTypeB)
	protoTypeList = append(protoTypeList, protoTypeB.Clone())

	for _, pt := range protoTypeList {
		fmt.Println(pt)
	}

	// 原型注册表实现
	registry := prototype.Registry{}
	registry.Register("ProtoTypeA", protoTypeA)
	registry.Register("ProtoTypeB", protoTypeB)

	ptA := registry.CreateProtoTypeById("ProtoTypeA")
	fmt.Println(ptA)

	ptB := registry.CreateProtoTypeById("ProtoTypeB")
	fmt.Println(ptB)
}
