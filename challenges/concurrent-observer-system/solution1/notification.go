package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type DeploymentNotification struct {
	event Event
}

func (d *DeploymentNotification) handle(event Event) {
	time.Sleep(time.Duration(rand.IntN(5)) * time.Second)
	fmt.Printf("Sending notification to %s's code owners \n", event.Service)
}

func (d *DeploymentNotification) getStatus() string {
	return d.event.Status
}
