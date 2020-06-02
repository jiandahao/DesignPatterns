package main

import (
	"fmt"
	"github.com/jiandahao/DesignPatterns/Behavioral_pattern/memento"
)

func main() {
	originator := memento.Originator{}
	originator.SetName("jian")
	snap := originator.CreateMemento("my-snapshot")
	{
		fmt.Println("====== Tesing get memento name and date ==========")
		fmt.Println(snap.GetMementoName())
		fmt.Println(snap.GetMementoDate().Format("2006-01-02 15:04:05"))
	}

	{
		fmt.Println("======= Tesing originator =========")
		originator.SetName("new name")
		fmt.Println(originator)
		originator.Restore(snap)
		fmt.Println(originator)
	}

	{
		// 将原发器的备忘和恢复托管给负责人（care taker）
		fmt.Println("======== Testing Caretaker ========")
		careTaker := memento.NewCaretaker(&originator)
		careTaker.ChangeName("new new name")
		fmt.Println(originator)

		careTaker.Undo()
		fmt.Println(originator)
	}
}
