package dao

import (
	"zhisheng-jing/loan-manager/global"
	. "zhisheng-jing/loan-manager/model"
)

type EmployeeDao struct{}

func NewEmployeeDao() *EmployeeDao {
	return &EmployeeDao{}
}

func (e EmployeeDao) GetEmployeeByID(id int) (*Employee,error) {
	employee := new(Employee)
	result:=global.DB.First(employee, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return employee, nil
}
func (e EmployeeDao) GetEmployees() ([]*Employee,error) {
	var employees []*Employee
	result := global.DB.Find(&employees)
	if result.Error !=nil {
		return nil,result.Error
	}

	return employees, nil
}

func (e EmployeeDao) GetEmployeeByCode(code string) (*Employee, error) {
	employee := new(Employee)

	result := global.DB.Where("employee_code=?",code).Find(&employee)

	if result.Error != nil {
		return  nil, result.Error
	}
	return employee, nil
}
