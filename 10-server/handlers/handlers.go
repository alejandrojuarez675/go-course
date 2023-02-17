package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/models"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	people := []models.Person{
		{Name: "ale", Age: 28},
	}

	peopleJson, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleJson)
}
