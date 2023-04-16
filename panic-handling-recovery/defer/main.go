package main

import "fmt"

func main() {
	// Defer the execution of Println() function
	// The order of execution of the defer statements will be LIFO (Last In First Out)
	defer fmt.Println("Three")
	defer fmt.Println("Four")
	defer fmt.Println("Five")

	fmt.Println("One")
	fmt.Println("Two")
}
