package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type DeploymentAudit struct {
	ch chan Event
}

func (d *DeploymentAudit) handle() {
	time.Sleep(time.Duration(rand.IntN(5)) * time.Second)
	for e := range d.ch {
		fmt.Printf("Saving deployment event for service %s\n", e.Service)
	}
}

func (d *DeploymentAudit) GetChan() chan Event { return d.ch }
