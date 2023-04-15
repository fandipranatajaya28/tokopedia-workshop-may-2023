package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

const (
	port = 8000
)

var (
	httpServer = &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}
)

func main() {
	// Let us see HTTP serve without graceful shutdown
	doServeHTTP()

	// Now let us implement graceful shutdown.
	// Can delete doServeHTTP to make playground simpler
	// doServeHTTPGraceful()
}

func doServeHTTP() {
	fmt.Println("Server is starting at port", port)
	err := httpServer.ListenAndServe()
	if err != nil {
		fmt.Println("error when ListenAndServe")
		return
	}

	fmt.Println("Process cleanup...") // This won't get called
}

func doServeHTTPGraceful() {
	// TODO: create context and its cancel function
	ctx, _ := context.WithCancel(context.Background())

	// TODO: setup SIGTERM listener
	go func() {
		// Listen for the termination signal

		// Block until termination signal received

		// Essentially the cancel() is broadcasted to all the goroutines that call .Done()
		// The returned context's Done channel is closed when the returned cancel function is called
	}()

	// setup errgroup with context so we can listen to its cancellation
	eg, egCtx := errgroup.WithContext(ctx)

	// setup HTTP listener
	eg.Go(func() error {
		fmt.Println("HTTP server listening on port", port)
		err := httpServer.ListenAndServe()
		fmt.Println("HTTP server finish listening on port", port)
		return err
	})

	// TODO: setup HTTP graceful shutdown
	eg.Go(func() error {
		<-egCtx.Done()
		return nil
	})

	// wait for errgroup
	if err := eg.Wait(); err != nil {
		fmt.Printf("Exit reason: %s \n", err)
	}

	fmt.Println("process cleanup...") // This should get called
}
