package main

import "fmt"

func main() {
	// buffered channel - limited capacity - queue-like
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}

	close(charChannel)

	// we can still loop over a closed channel
	for result := range charChannel {
		fmt.Println(result)
	}

}
