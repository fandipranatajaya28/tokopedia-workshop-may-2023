package main

import (
	"fmt"
	"time"

	"github.com/fandipranatajaya28/tokopedia-workshop-may-2023/panic-handler/wrapper"
)

func a() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var i *int
	fmt.Println(*i)
}

func b() {
	i := 10
	go wrapper.PanicHandleGoRoutine(func() {
		fmt.Println(i)
		panic("panicking")
	})
}

func main() {
	a()
	b()
	time.Sleep(10 * time.Second)
	fmt.Println("returns normally")
}
