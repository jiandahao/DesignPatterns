package main

import "github.com/jiandahao/DesignPatterns/Structural_patterns/bridge"

func main() {
	{
		impl := &bridge.TV{}
		abst := bridge.NewRemoteAbstraction(impl)
		abst.Operate()
	}

	{
		impl := &bridge.AirCondition{}
		abst := bridge.NewRemoteAbstraction(impl)
		abst.Operate()
	}

	{
		impl := &bridge.TV{}
		abst := bridge.NewExtendRemoteAbstraction(impl)
		abst.Operate()
		abst.MultiOperate()
	}
}
