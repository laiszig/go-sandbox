package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type DeploymentNotification struct {
	ch chan Event
}

func (d *DeploymentNotification) handle() {
	time.Sleep(time.Duration(rand.IntN(5)) * time.Second)
	for e := range d.ch {
		fmt.Printf("Sending notification to %s's code owners \n", e.Service)
	}
}

func (d *DeploymentNotification) GetChan() chan Event { return d.ch }
