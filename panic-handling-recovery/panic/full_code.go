package main

import "fmt"

func FullCode_firstFunction() {
	fmt.Println("First function called")
	FullCode_secondFunction()
	fmt.Println("First function finished") // This will not get called
}

func FullCode_secondFunction() {
	fmt.Println("Second function called")
	panic("Panic happens")                  // Go library for panic
	fmt.Println("Second function finished") // This will not get called
}

func FullCode_doPanic() {
	fmt.Println("Panic example in Go")
	FullCode_firstFunction()
	fmt.Println("All process finished") // This will not get called
}
