package model

//员工
type Employee struct {
	BaseEmployee
	EmployeeCode   string       `json:"employee_code"`                                   //工号
	Email          string	`json:"email"`
	Name           string               `json:"name"`                                              //姓名
}

func (e Employee) TableName() string {
		return "employee"
}