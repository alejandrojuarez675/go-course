package main

import (
	"log"
	"net/http"
	"packageserver/handlers"
)

const PORT = "3000"

func main() {
	// router
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlers.PingHandler)

	// create server
	server := &http.Server{
		Addr:    "localhost:" + PORT,
		Handler: mux,
	}
	log.Println("Start server on port " + PORT)
	log.Fatal(server.ListenAndServe())
}
