package main

import (
	"log"
	"net/http"
	"server/handlers"

	"github.com/gorilla/mux"
)

const PORT = "3000"

func main() {
	// router
	mux := mux.NewRouter()
	mux.HandleFunc("/ping", handlers.PingHandler).Methods("GET")
	mux.HandleFunc("/people/{id:[0-9]+}", handlers.GetPeopleById).Methods("GET")
	mux.HandleFunc("/people", handlers.GetAllPeople).Methods("GET")

	// create server
	server := &http.Server{
		Addr:    "localhost:" + PORT,
		Handler: mux,
	}
	log.Println("Start server on port " + PORT)
	log.Fatal(server.ListenAndServe())
}
