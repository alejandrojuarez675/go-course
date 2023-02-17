package handlers

import (
	"net/http"
	"server/dao"

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
