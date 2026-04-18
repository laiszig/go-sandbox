package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			select {
			default:
				fmt.Println("Doing Work...")
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
