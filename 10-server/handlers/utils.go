package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
)

func returnJSON(w http.ResponseWriter, data any, status int) {

	if status >= 400 {
		data = models.ErrorResponse{StatusCode: status, Error: data}
	}

	dataJson, err := json.Marshal(data)
	if err != nil {
		returnJSON(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dataJson)
}
