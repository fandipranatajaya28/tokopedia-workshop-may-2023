package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func FullCode_doProcessGraceful() {
	ctx, cancel := context.WithCancel(context.Background())

	// do setup SIGTERM listener
	go func() {
		// Listen for the termination signal
		stop := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		// Block until termination signal received
		<-stop
		// Essentially the cancel() is broadcasted to all the goroutines that call .Done()
		// The returned context's Done channel is closed when the returned cancel function is called
		cancel()
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done(): // Block until cancel() is called
				fmt.Println("Break the first loop")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Hello in the first loop")
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done(): // Block until cancel() is called
				fmt.Println("Break the second loop")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Hello in the second loop")
			}
		}
	}()

	// Wait for ongoing process to finish
	wg.Wait()
	fmt.Println("Process cleanup...") // This will be called
}
