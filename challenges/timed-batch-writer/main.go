package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const bufferSize = 100
const timerDuration = 10 * time.Second

func CreateShutdownContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-ch
		fmt.Println("Received termination signal, shutting down gracefully")
		cancel()
	}()

	return ctx
}

func main() {
	ctx := CreateShutdownContext()

	var g sync.WaitGroup
	savingCh := make(chan int, 0)
	g.Go(func() {
		receive(savingCh)
	})

	g.Go(func() {
		s := make([]int, 0)
		timer := time.NewTimer(timerDuration)
		defer timer.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context cancelled")
				return
			case <-timer.C:
				if len(s) > 0 {
					save(s)
				}
				s = make([]int, 0)
				timer.Reset(timerDuration)
			case num := <-savingCh:
				s = append(s, num)
				if len(s) >= bufferSize {
					save(s)
					s = make([]int, 0)
					timer.Reset(timerDuration)
				}
			}
		}

	})

	g.Wait()
}

func receive(ch chan int) {
	for {
		time.Sleep(time.Duration(rand.IntN(200)) * time.Millisecond)
		ch <- rand.IntN(999_999_999)
	}
}

func save(numbers []int) {
	fmt.Println("Saving numbers: ", len(numbers))
}
