package main

import (
	"fmt"
	"github.com/jiandahao/DesignPatterns/Creational_patterns/singleton"
)

func main() {
	dbSingleton := singleton.DatabaseSingleton{}
	dbA := dbSingleton.GetInstance()
	dbB := dbSingleton.GetInstance()
	fmt.Println(dbA)
	fmt.Println(dbB)
	fmt.Printf("dbA: %x\n", *dbA)
	fmt.Printf("dbB: %x\n", *dbB)
}
