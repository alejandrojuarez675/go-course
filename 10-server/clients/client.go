package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/models/employee"
	"server/models/externalClientError"
)

type ResponseType interface {
	employee.EmployeeResponse | employee.EmployeeListResponse
}

func Get[T ResponseType](url string) (T, externalClientError.ExternalClientError) {
	resp, err := http.Get(url)

	var bodyResponse T
	if err != nil {
		return bodyResponse, externalClientError.GetErrorConnection()
	} else {

		fmt.Println("Service "+url+" response", resp.Status)
		if resp.StatusCode != 200 {
			return bodyResponse, externalClientError.GetExpectationFailed(resp.Status)
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body) // response body is []byte

		if err := json.Unmarshal(body, &bodyResponse); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON")
			return bodyResponse, externalClientError.GetUnprocessableEntityError()
		}
		return bodyResponse, externalClientError.ExternalClientError{}
	}
}
