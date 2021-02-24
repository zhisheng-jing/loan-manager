package service

import (
	"zhisheng-jing/loan-manager/dao"
	"zhisheng-jing/loan-manager/model"
)

type EmployeeService struct {
	
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{}
}

func (e EmployeeService) GetEmployeeByID(id int) (*model.Employee,error) {
	return dao.NewEmployeeDao().GetEmployeeByID(id)
}

func (e EmployeeService) GetEmployees() ([]*model.Employee, error) {
	return dao.NewEmployeeDao().GetEmployees()
}

func (e EmployeeService) GetEmployeeByCode(code string) (*model.Employee, error) {
	return dao.NewEmployeeDao().GetEmployeeByCode(code)
}

