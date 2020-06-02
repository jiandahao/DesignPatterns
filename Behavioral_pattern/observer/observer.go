package observer

import (
	"fmt"
	"reflect"
)

type Publisher struct {
	subscribers []Subscriber
}

func (publisher *Publisher) AddSubscriber(subscriber Subscriber) {
	publisher.subscribers = append(publisher.subscribers, subscriber)
}

func (publisher *Publisher) RemoveSubscriber(subscriber Subscriber) {
	for index, sub := range publisher.subscribers {
		if reflect.DeepEqual(sub, subscriber) {
			publisher.subscribers = append(publisher.subscribers[:index], publisher.subscribers[index+1:]...)
		}
	}
}

func (publisher *Publisher) NotifySubscribers() {
	for _, sub := range publisher.subscribers {
		sub.Update()
	}
}


type Subscriber interface {
	Update()
}

type ConcreteSubscriberA struct {}

func (sub *ConcreteSubscriberA) Update() {
	fmt.Println("Concrete subscriber A update")
}

type ConcreteSubscriberB struct {}

func (sub *ConcreteSubscriberB) Update() {
	fmt.Println("Concrete subscriber B update")
}