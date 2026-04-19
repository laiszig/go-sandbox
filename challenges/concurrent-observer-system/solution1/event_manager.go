package main

import "sync"

type EventManager struct {
	observers []Observer
}

func (d *EventManager) register(observer Observer) {
	d.observers = append(d.observers, observer)
}

func (d *EventManager) notifyAll(event Event) {
	wg := sync.WaitGroup{}
	for _, observer := range d.observers {
		wg.Add(1)
		go func(observer Observer, event Event) {
			defer wg.Done()
			observer.handle(event)
		}(observer, event)
	}

	// api responds after the longest running go routine
	wg.Wait() // if commented, API can respond immediately
}
