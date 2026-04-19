package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type DeploymentLogger struct {
	ch chan Event
}

func (d *DeploymentLogger) handle() {
	time.Sleep(time.Duration(rand.IntN(5)) * time.Second)
	for e := range d.ch {
		fmt.Printf("Logging deployment for service %s\n", e.Service)
	}
}

func (d *DeploymentLogger) GetChan() chan Event { return d.ch }
