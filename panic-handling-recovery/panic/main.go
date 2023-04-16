package main

import "fmt"

func firstFunction() {
	fmt.Println("First function called")
	secondFunction()
	fmt.Println("First function finished") // This will not get called
}

func secondFunction() {
	fmt.Println("Second function called")

	// TODO: execute panic to trigger termination of the application

	fmt.Println("Second function finished") // This should not get called
}

func doPanic() {
	fmt.Println("Panic example in Go")
	firstFunction()
	fmt.Println("All process finished") // This should not get called
}

func main() {
	doPanic()
}
