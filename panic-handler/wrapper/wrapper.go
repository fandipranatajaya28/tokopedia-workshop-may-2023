package wrapper

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

type CommandFunc func()

func PanicHandleHTTP(command http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				errString, _ := err.(string)
				fmt.Println(errString)
				http.Error(w, errString, http.StatusInternalServerError)
				debug.PrintStack()
			}
		}()

		command(w, r)
	}
}

func PanicHandleGoRoutine(command CommandFunc) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			log.Println(string(debug.Stack()))
		}
	}()

	command()
}
