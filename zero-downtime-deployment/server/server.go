package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/fandipranatajaya28/tokopedia-workshop-may-2023/zero-downtime-deployment/zero_interruption"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("initiate new server")

	// assume it took time to connect to database
	log.Println("connecting to database...")
	time.Sleep(1 * time.Second)
	log.Println("database connected")

	// assume it took time to build cache
	log.Println("building cache...")
	time.Sleep(1 * time.Second)
	log.Println("cache built")

	// assume it took time to make sure upstream ready
	log.Println("ping upstream service...")
	time.Sleep(1 * time.Second)
	log.Println("upstream ready")

	handlers := mux.NewRouter()
	handlers.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "Make it Happen\n")
		// io.WriteString(res, "Make it Better\n")
	})

	srv := http.Server{
		Handler:      handlers,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("initiate server finished")

	zero_interruption.ServeHTTP("8000", srv)
}
