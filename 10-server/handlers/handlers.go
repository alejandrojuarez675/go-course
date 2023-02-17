package handlers

import (
	"net/http"
	"server/clients/dummyClient"
	"server/dao"
	"server/models/externalClientError"

	"github.com/gorilla/mux"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	returnJSON(w, "pong", http.StatusOK)
}

func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	people := dao.GetAllPeople()
	returnJSON(w, people, http.StatusOK)
}

func GetPeopleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := vars["id"]

	person, err := dao.GetPersonById(userId)
	if err != nil {
		returnJSON(w, "Resource not found", http.StatusNotFound)
	} else {
		returnJSON(w, person, http.StatusOK)
	}
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := dummyClient.GetEmployees()
	if err != (externalClientError.ExternalClientError{}) {
		returnJSON(w, err.Message, err.Status)
	} else {
		returnJSON(w, employees, http.StatusOK)
	}
}
