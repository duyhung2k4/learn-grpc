package main

import (
	"app/grpc"
	"app/router"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		grpc.ServerGRPC()
		wg.Done()
	}()

	go func() {
		server := http.Server{
			Addr:         "localhost:8080",
			Handler:      router.Router(),
			ReadTimeout:  time.Second * 30,
			WriteTimeout: time.Second * 30,
		}

		log.Fatalln(server.ListenAndServe())

		wg.Done()
	}()

	wg.Wait()
}
