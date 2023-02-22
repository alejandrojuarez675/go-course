package dummyClient

import (
	"server/clients"
	"server/models/employee"
	"server/models/externalClientError"
)

func GetEmployees() ([]employee.Employee, externalClientError.ExternalClientError) {
	url := "https://dummy.restapiexample.com/api/v1/employees"

	resp, err := clients.Get[employee.EmployeeListResponse](url)
	return resp.Data, err
}

func GetEmployeeById(id string) (employee.Employee, externalClientError.ExternalClientError) {
	url := "https://dummy.restapiexample.com/api/v1/employee/" + id

	resp, err := clients.Get[employee.EmployeeResponse](url)
	return resp.Data, err
}
