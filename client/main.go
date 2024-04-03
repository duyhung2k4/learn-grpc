package main

import (
	"app/router"
	"log"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr:         "localhost:8081",
		Handler:      router.Router(),
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	log.Fatalln(server.ListenAndServe())
}
