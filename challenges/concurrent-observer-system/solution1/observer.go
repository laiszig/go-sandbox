package main

type Observer interface {
	handle(event Event)
}
