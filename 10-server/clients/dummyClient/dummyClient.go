package dummyClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/models"
)

func GetEmployees() ([]models.Employee, models.ExternalClientError) {
	resp, err := http.Get("https://dummy.restapiexample.com/api/v1/employees")
	if err != nil {
		errorToReturn := models.ExternalClientError{
			Status:  http.StatusExpectationFailed,
			Message: "Problem in connection with external service",
		}
		return []models.Employee{}, errorToReturn
	} else {

		if resp.StatusCode != 200 {
			fmt.Println("Service response", resp.StatusCode)
			errorToReturn := models.ExternalClientError{
				Status:  http.StatusExpectationFailed,
				Message: "Error on external client. Service response with " + resp.Status,
			}
			return []models.Employee{}, errorToReturn
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body) // response body is []byte

		var bodyResponse models.EmployeeResponse
		if err := json.Unmarshal(body, &bodyResponse); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON")
			errorToReturn := models.ExternalClientError{
				Status:  http.StatusUnprocessableEntity,
				Message: "Can not unmarshal JSON",
			}
			return []models.Employee{}, errorToReturn
		}
		return bodyResponse.Data, models.ExternalClientError{}
	}
}
