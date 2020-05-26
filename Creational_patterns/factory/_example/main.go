package main

import (
	"github.com/jiandahao/DesignPatterns/Creational_patterns/factory"
	"math/rand"
	"time"
)

func main() {
	var dialog factory.DialogFactory
	rand.Seed(time.Now().Unix())
	osType := []string{"linux", "windows"}
	typ := osType[rand.Intn(10)%2]
	if "linux" == typ {
		dialog = &factory.LinuxDialogFactory{}
	} else if "windows" == typ {
		dialog = &factory.WindowDialogFactory{}
	} else {
		return
	}

	button := dialog.CreateButton()
	button.OnClick()
	button.Render()
}
