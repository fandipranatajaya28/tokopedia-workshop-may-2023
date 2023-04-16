package main

import "fmt"

func FullCode_doDefer() {
	// Defer the execution of Println() function
	// The order of execution of the defer statements will be LIFO (Last In First Out)
	defer fmt.Println("Five")
	defer fmt.Println("Four")
	defer fmt.Println("Three")

	fmt.Println("One")
	fmt.Println("Two")
}
