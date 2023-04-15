package main

import "fmt"

func firstFunction() {
	// Defer the process of recovery
	defer func() {
		// Recover from panic to stop termination of the application
		if r := recover(); r != nil {
			fmt.Printf("Panic message: %+v\n", r)
			fmt.Println("First function recovered from the panic")
		}
	}()
	fmt.Println("First function called")
	secondFunction()
	fmt.Println("First function finished") // This will not get called
}

func secondFunction() {
	fmt.Println("Second function called")
	panic("Panic happens")                  // Go library for panic
	fmt.Println("Second function finished") // This will not get called
}

func main() {
	fmt.Println("Panic example in Go")
	firstFunction()
	fmt.Println("Function main done") // This will get called
}
