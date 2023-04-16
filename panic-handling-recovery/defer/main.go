package main

import "fmt"

func doDefer() {
	// Defer the execution of Println() function
	// The order of execution of the defer statements will be LIFO (Last In First Out)

	// TODO: print One, Two, Three, Four, Five sequentially using defer

	fmt.Println("One")
	fmt.Println("Two")
}

func main() {
	doDefer()
}
