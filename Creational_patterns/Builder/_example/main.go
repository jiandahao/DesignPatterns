package main

import (
	"fmt"
	builder "github.com/jiandahao/DesignPatterns/Creational_patterns/Builder"
)

func main() {
	// create a builder
	aConcreteBuilder := builder.NewConcreteBuilder()

	// create a director and set build
	aDirector := builder.NewDirector(aConcreteBuilder)

	// construct product
	aDirector.Construct()

	// get results
	res := aConcreteBuilder.GetResult()
	fmt.Println(res)
}
