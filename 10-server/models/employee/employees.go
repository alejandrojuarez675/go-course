package employee

type EmployeeListResponse struct {
	Status  string
	Data    []Employee
	Message string
}

type EmployeeResponse struct {
	Status  string
	Data    Employee
	Message string
}

type Employee struct {
	Id             int32  `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary int64  `json:"employee_salary"`
	EmployeeAge    int8   `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}
