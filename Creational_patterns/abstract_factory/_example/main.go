package main

import (
	"github.com/jiandahao/DesignPatterns/Creational_patterns/abstract_factory"
	"math/rand"
	"time"
)

type Config struct {
	UIType string
}

func getFactoryByConfig(config Config) abstract_factory.GUIFactory {
	switch config.UIType {
	case "linux":
		return &abstract_factory.LinuxGUIFactory{}
	case "windows":
		return &abstract_factory.WindowsGUIFactory{}
	default:
		return nil
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	uiType := []string{"linux", "windows"}
	cfg := Config{UIType: uiType[rand.Intn(2)]}

	factory := getFactoryByConfig(cfg)

	button := factory.CreateButton()
	checkBox := factory.CreateCheckBox()

	button.OnClick()
	checkBox.Check()
}
