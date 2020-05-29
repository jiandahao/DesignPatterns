package main

import (
	"fmt"
	"github.com/jiandahao/DesignPatterns/Structural_patterns/decorator"
)

func main() {
	{
		mem := decorator.NewBuffer()
		mem.WriteData([]byte("test"))
		fmt.Println(string(mem.ReadData()))

		file := decorator.NewFile()
		file.WriteData([]byte("test"))
		file.ReadData()
	}

	{

		fmt.Println("=======Testing decorator===========")
		mem := decorator.NewBuffer()
		encryDec := decorator.NewEncryptionDecorator(mem)
		validatorDec := decorator.NewValidateDecorator(encryDec)
		validatorDec.WriteData([]byte("TEXT"))
		fmt.Println(string(validatorDec.ReadData()))
		//encryDec.WriteData([]byte("TEST"))
		//fmt.Println(string(encryDec.ReadData()))
	}

}
