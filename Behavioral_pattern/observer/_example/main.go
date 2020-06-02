package main

import "github.com/jiandahao/DesignPatterns/Behavioral_pattern/observer"

func main() {
	subA := &observer.ConcreteSubscriberA{}
	subB := &observer.ConcreteSubscriberB{}

	publisher := observer.Publisher{}
	publisher.AddSubscriber(subA)
	publisher.AddSubscriber(subB)

	publisher.NotifySubscribers()
}
