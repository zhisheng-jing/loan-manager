package model

//员工
type Employee struct {
	ID             int  `json:"id"`
	EmployeeCode   string       `json:"employee_code"`                                   //工号
	Email          string	`json:"email"`
	Name           string               `json:"name"`                                              //姓名
}

func (emply Employee) TableName() string {
		return "employee"
}