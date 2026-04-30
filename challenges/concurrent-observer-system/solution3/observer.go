package main

type Observer interface {
	handle()
	GetChan() chan Event
}
