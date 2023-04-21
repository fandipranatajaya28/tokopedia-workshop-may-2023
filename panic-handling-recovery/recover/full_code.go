package main

import (
	"fmt"
	"runtime/debug"
)

func FullCode_firstFunction() {
	// Defer the process of recovery
	defer func() {
		// Recover from panic to stop termination of the application
		if r := recover(); r != nil {
			fmt.Printf("Panic message: %+v\n", r)
			fmt.Println("First function recovered from the panic")
			// use debug.PrintStack() if you want to trace the panic and print it
			debug.PrintStack()
		}
	}()
	fmt.Println("First function called")
	FullCode_secondFunction()
	fmt.Println("First function finished") // This will not get called
}

func FullCode_secondFunction() {
	fmt.Println("Second function called")
	panic("Panic happens")                  // Go library for panic
	fmt.Println("Second function finished") // This will not get called
}

func FullCode_doRecover() {
	fmt.Println("Panic example in Go")
	FullCode_firstFunction()
	fmt.Println("All process finished") // This will get called
}
