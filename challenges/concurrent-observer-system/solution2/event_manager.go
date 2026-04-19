package main

import (
	"fmt"
	"sync"
)

type EventManager struct {
	observers []Observer
}

func (d *EventManager) register(observer Observer) {
	d.observers = append(d.observers, observer)
	go observer.handle()
}

func (d *EventManager) notifyAll(event Event) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer wg.Done()

	for _, observer := range d.observers {
		select {
		case observer.GetChan() <- event:
			fmt.Printf("Observed event for service %s\n", event.Service)
		}
	}
	wg.Wait() // if commented, API can respond immediately
}
