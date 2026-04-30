package main

import (
	"fmt"
)

type EventManager struct {
	observers []Observer
}

func (d *EventManager) register(observer Observer) {
	d.observers = append(d.observers, observer)
	go observer.handle()
}

func (d *EventManager) notifyAll(event Event) {
	for _, observer := range d.observers {
		select {
		case observer.GetChan() <- event:
			fmt.Printf("Observed event for service %s\n", event.Service)
		}
	}
}
