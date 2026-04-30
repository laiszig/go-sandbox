package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//func createGracefulShutdownCtx(ctx context.Context) context.Context {
//	newCtx, cancel := context.WithCancel(ctx)
//	signalChan := make(chan os.Signal, 1)
//
//	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
//
//	go func() {
//		select {
//		case <-signalChan:
//			cancel()
//		}
//	}()
//	return newCtx
//}

// with timeout
func createGracefulShutdownCtx(ctx context.Context) context.Context {
	newCtx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-signalChan:
			time.Sleep(5 * time.Second)
			cancel()
		}
	}()
	return newCtx
}

func main() {

	ctxBackground := context.Background()
	ctxTimeout, _ := context.WithTimeout(ctxBackground, 5*time.Second)
	ctx := createGracefulShutdownCtx(ctxTimeout)

	select {
	case <-ctx.Done():
		fmt.Println("shutting down")
	}
}
