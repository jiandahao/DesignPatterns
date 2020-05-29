package adapter

import (
	"strconv"
)

// The Target defines the domain-specific interface used by the client code.
type Target interface {
	Request() string
}

type NormalTarget struct{}

func (t *NormalTarget) Request() string {
	return "1234567"
}

/*
 * The Adaptee contains some useful behavior, but its interface is incompatible
 * with the existing client code. The Adaptee needs some adaptation before the
 * client code can use it.
 */
type Adaptee struct{}

func (a *Adaptee) SpecialRequest() int {
	return 1234567
}

/**
 * The Adapter makes the Adaptee's interface compatible with the Target's
 * interface.
 */
type Adapter struct {
	Adaptee Adaptee
}

func NewAdapter(adaptee Adaptee) *Adapter {
	return &Adapter{Adaptee: adaptee}
}
func (a *Adapter) Request() string {
	res := a.Adaptee.SpecialRequest()
	// convert number into string
	return strconv.Itoa(res)
}
