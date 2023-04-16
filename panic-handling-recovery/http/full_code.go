package main

import (
	"fmt"
	"net/http"
)

func FullCode_panicHandleHTTP(command http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Defer the process of recovery
		defer func() {
			// Recover from panic to stop termination of the application
			if err := recover(); err != nil {
				fmt.Printf("Panic message: %+v\n", err)
				fmt.Println("Function recovered from the panic")
			}
		}()

		// Execute HTTP function that has been wrapped
		command(w, r)
	}
}

func FullCode_registerRoutes(server *http.Server) {
	// Create endpoint to test panic process and call HTTP wrapper function to wrap our process
	http.Handle("/test", FullCode_panicHandleHTTP(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello\n")) // Write message to the client
		panic("Panic happens")     // Go library for panic
	}))
	server.Handler = http.DefaultServeMux
}

func FullCode_doHTTPPanicRecovery() {
	port := 8000
	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	// Register our HTTP endpoint
	FullCode_registerRoutes(httpServer)

	fmt.Println("HTTP server listening on port", port)
	err := httpServer.ListenAndServe()
	if err != nil {
		fmt.Println("error when ListenAndServe")
		return
	}
}
