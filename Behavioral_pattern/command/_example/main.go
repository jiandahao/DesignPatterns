package main

import "github.com/jiandahao/DesignPatterns/Behavioral_pattern/command"

func main() {
	invoker := command.NewInvoker() // in real world, invoker could be a button
	invoker.SetCommand("Ctrl+C", command.NewCopyCommand())
	invoker.SetCommand("Ctrl+V", command.NewPasteCommand())

	invoker.Execute("Ctrl+C")
	invoker.Execute("Ctrl+V")
}
