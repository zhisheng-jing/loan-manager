package dao

import (
	"errors"
	"zhisheng-jing/loan-manager/global"
	. "zhisheng-jing/loan-manager/model"
)

type EmployeeDao struct{}

func NewEmployeeDao() *EmployeeDao {
	return &EmployeeDao{}
}

func (e EmployeeDao) GetEmployeeByID(id int) (Employee,error) {
	var employee Employee
	result:=global.DB.First(&employee, id)
	if result.RowsAffected == 0 {
		return employee, result.Error
	}
	return employee, nil
}
func (e EmployeeDao) GetEmployees() ([]Employee,error) {
	var employees []Employee
	result := global.DB.Find(&employees)
	if result.RowsAffected == 0 {
		return employees, result.Error
	}
	return employees, nil
}

func (e EmployeeDao) GetEmployeeByCode(code string) (Employee, error) {
	var employee Employee
	result := global.DB.Where("employee_code=?",code).Find(&employee)
	if result.RowsAffected == 0{
		return employee,errors.New("找不到")
	}

	return employee, nil
}
