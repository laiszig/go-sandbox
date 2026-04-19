package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type DeploymentAudit struct {
	event Event
}

func (d *DeploymentAudit) handle(event Event) {
	time.Sleep(time.Duration(rand.IntN(5)) * time.Second)
	fmt.Printf("Saving deployment event for service %s\n", event.Service)
}
