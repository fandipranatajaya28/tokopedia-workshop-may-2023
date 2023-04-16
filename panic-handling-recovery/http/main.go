package main

import (
	"fmt"
	"net/http"
)

func panicHandleHTTP(command http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Defer the process of recovery
		defer func() {
			// Recover from panic to stop termination of the application

			// TODO: setup recover function to recover from a panic
		}()

		// Execute HTTP function that has been wrapped
		command(w, r)
	}
}

func registerRoutes(server *http.Server) {
	// Create endpoint to test panic process and call HTTP wrapper function to wrap our process
	// TODO: call HTTP wrapper function using http.Handle()
	// TODO: write message to the client and execute panic to trigger termination of the application

	server.Handler = http.DefaultServeMux
}

func doHTTPPanicRecovery() {
	port := 8000
	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	// Register our HTTP endpoint
	registerRoutes(httpServer)

	fmt.Println("HTTP server listening on port", port)
	err := httpServer.ListenAndServe()
	if err != nil {
		fmt.Println("error when ListenAndServe")
		return
	}
}

func main() {
	doHTTPPanicRecovery()
}
