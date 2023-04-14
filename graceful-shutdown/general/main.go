package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// let us see process without graceful will behave
	doProcess()

	// then, let us work on graceful implementation
	// can delete doProcess function to make playground simpler
	//doProcessGraceful()
}

func doProcess() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			time.Sleep(1 * time.Second)
			fmt.Println("Hello in the first loop")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			time.Sleep(1 * time.Second)
			fmt.Println("Hello in the second loop")
		}
	}()

	wg.Wait()
	fmt.Println("Process cleanup...") // this won't get called
}

func doProcessGraceful() {
	// TODO: setup context and its cancel function

	// TODO: setup SIGTERM listener
	go func() {
		// Listen for the termination signal

		// Block until termination signal received

		// Essentially the cancel() is broadcasted to all the goroutines that call .Done()
		// The returned context's Done channel is closed when the returned cancel function is called
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		// TODO: convert into select case syntax and listen to context cancellateion
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("Hello in the first loop")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		// TODO: convert into select case syntax and listen to context cancellateion
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("Hello in the second loop")
		}
	}()

	// Wait for ongoing process to finish
	wg.Wait()
	fmt.Println("Process cleanup...") // this should be called
}
