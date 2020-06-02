package main

import "github.com/jiandahao/DesignPatterns/Behavioral_pattern/mediator"

func main() {

	m := &mediator.ConcreteMediator{}

	button := mediator.NewButton(m)
	checkBox := mediator.NewCheckBox(m)

	m.Button = button
	m.CheckBox = checkBox

	button.OnClick()

	checkBox.OnClick()
}
