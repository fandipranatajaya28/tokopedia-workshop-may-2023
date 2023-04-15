package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func FullCode_doServeHTTPGraceful() {
	ctx, cancel := context.WithCancel(context.Background())

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

	port := 8000
	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		fmt.Println("HTTP server listening on port", port)
		err := httpServer.ListenAndServe()
		fmt.Println("HTTP server finish listening on port", port)
		return err
	})

	eg.Go(func() error {
		// Block until cancel() is called
		<-egCtx.Done()
		fmt.Println("HTTP server start graceful shutdown on port", port)
		err := httpServer.Shutdown(context.Background()) // Go library for HTTP server graceful shutdown
		fmt.Println("HTTP server finish graceful shutdown on port", port)
		return err
	})

	// Wait for ongoing process to finish
	if err := eg.Wait(); err != nil {
		fmt.Printf("Exit reason: %s \n", err)
	}

	fmt.Println("Process cleanup...") // This will be called
}
