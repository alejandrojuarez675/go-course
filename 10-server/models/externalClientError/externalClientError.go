package externalClientError

import (
	"net/http"
)

type ExternalClientError struct {
	Status  int
	Message string
}

func GetErrorConnection() ExternalClientError {
	return ExternalClientError{
		Status:  http.StatusExpectationFailed,
		Message: "Problem in connection with external service",
	}
}

func GetExpectationFailed(status string) ExternalClientError {
	return ExternalClientError{
		Status:  http.StatusExpectationFailed,
		Message: "Error on external client. Service response with " + status,
	}
}

func GetUnprocessableEntityError() ExternalClientError {
	return ExternalClientError{
		Status:  http.StatusUnprocessableEntity,
		Message: "Can not unmarshal JSON",
	}
}
