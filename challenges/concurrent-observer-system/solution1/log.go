package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type DeploymentLogger struct {
	event Event
}

func (d *DeploymentLogger) handle(event Event) {
	time.Sleep(time.Duration(rand.IntN(5)) * time.Second)
	fmt.Printf("Logging deployment for service %s\n", event.Service)
}
