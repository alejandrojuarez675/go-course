package dummyClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/models/employee"
	"server/models/externalClientError"
)

func GetEmployees() ([]employee.Employee, externalClientError.ExternalClientError) {
	resp, err := http.Get("https://dummy.restapiexample.com/api/v1/employees")
	if err != nil {
		return []employee.Employee{}, externalClientError.GetErrorConnection()
	} else {

		if resp.StatusCode != 200 {
			fmt.Println("Service response", resp.StatusCode)
			return []employee.Employee{}, externalClientError.GetExpectationFailed(resp.Status)
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body) // response body is []byte

		var bodyResponse employee.EmployeeResponse
		if err := json.Unmarshal(body, &bodyResponse); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON")
			return []employee.Employee{}, externalClientError.GetUnprocessableEntityError()
		}
		return bodyResponse.Data, externalClientError.ExternalClientError{}
	}
}
