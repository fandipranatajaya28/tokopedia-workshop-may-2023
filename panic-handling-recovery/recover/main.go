package main

import "fmt"

func firstFunction() {
	// Defer the process of recovery
	defer func() {
		// Recover from panic to stop termination of the application

		// TODO: setup recover function to recover from a panic
	}()
	fmt.Println("First function called")
	secondFunction()
	fmt.Println("First function finished") // This should not get called
}

func secondFunction() {
	fmt.Println("Second function called")
	panic("Panic happens")                  // Go library for panic
	fmt.Println("Second function finished") // This should not get called
}

func doRecover() {
	fmt.Println("Panic example in Go")
	firstFunction()
	fmt.Println("All process finished") // This should get called
}

func main() {
	doRecover()
}
