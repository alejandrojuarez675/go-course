package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "3000"

func main() {
	// router
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "pong") })

	// create server
	server := &http.Server{
		Addr:    "localhost:" + PORT,
		Handler: mux,
	}
	log.Println("Start server on port " + PORT)
	log.Fatal(server.ListenAndServe())
}
