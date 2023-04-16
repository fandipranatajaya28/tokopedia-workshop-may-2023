package wrapper

import (
	"fmt"
	"net/http"
)

type CommandFunc func()

// Wrapper function for HTTP panic handling
func PanicHandleHTTP(command http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Defer the process of recovery
		defer func() {
			// Recover from panic to stop termination of the application
			if err := recover(); err != nil {
				fmt.Printf("Panic message: %+v\n", err)
				fmt.Println("Function recovered from the panic")
				// use debug.PrintStack() if you want to trace the panic and print it
				// debug.PrintStack()
			}
		}()

		// Execute HTTP function that has been wrapped
		command(w, r)
	}
}

// Wrapper function for general panic handling
func PanicHandleGoRoutine(command CommandFunc) {
	// Defer the process of recovery
	defer func() {
		// Recover from panic to stop termination of the application
		if err := recover(); err != nil {
			fmt.Printf("Panic message: %+v\n", err)
			fmt.Println("Function recovered from the panic")
			// use debug.Stack() if you want to trace the panic
			// log.Println(string(debug.Stack()))
		}
	}()

	// Execute function that has been wrapped
	command()
}
