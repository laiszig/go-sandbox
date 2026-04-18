package main

import "fmt"

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}() // anonymous function is invoked

	go func() {
		anotherChannel <- "data2"
	}()

	//msg := <-myChannel // blocking line - main will wait for it - this is the join point
	// the main function will wait for the ch to close
	// or a message to be received

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgmFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgmFromAnotherChannel)
	}

	//fmt.Println(msg)
}
