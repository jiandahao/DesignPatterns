package bridge

// 将一个大类或一系列紧密相关的类拆分为抽象和实现两个独立的层次结构， 从而能在开发时分别使用。
import "fmt"

// 遥控器抽象部分
type RemoteAbstraction struct {
	device DeviceImplementation
}

func NewRemoteAbstraction(device DeviceImplementation) *RemoteAbstraction {
	return &RemoteAbstraction{device: device}
}

func (r *RemoteAbstraction) Operate() {
	r.device.TurnOn()
	r.device.TurnOff()
}

type ExtendRemoteAbstraction struct {
	*RemoteAbstraction
}

func NewExtendRemoteAbstraction(device DeviceImplementation) *ExtendRemoteAbstraction {
	return &ExtendRemoteAbstraction{NewRemoteAbstraction(device)}

}

func (r *ExtendRemoteAbstraction) MultiOperate() {
	// turn on twice and off once
	r.device.TurnOn()
	r.device.TurnOn()
	r.device.TurnOff()
}

// 设备实现部分，这相当于就是这个桥，抽象部分只需要操作这个一个接口实现，就可以
// 在运行时根据包含的实现部分做对应的事情。
type DeviceImplementation interface {
	TurnOn()
	TurnOff()
}

type TV struct{}

func (tv *TV) TurnOn() {
	fmt.Println("turn on tv")
}

func (tv *TV) TurnOff() {
	fmt.Println("turn off tv")
}

type AirCondition struct{}

func (a *AirCondition) TurnOn() {
	fmt.Println("turn on air condition")
}

func (a *AirCondition) TurnOff() {
	fmt.Println("turn off air condition")
}
