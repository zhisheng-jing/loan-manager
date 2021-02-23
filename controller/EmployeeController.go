package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zhisheng-jing/loan-manager/service"
)

type EmployeeController struct{}

func NewEmployeeController() *EmployeeController {
	return &EmployeeController{}
}

func (e *EmployeeController) GetEmployeeByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	employee, err := service.NewEmployeeService().GetEmployeeByID(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, &employee)
	return
}
func (e *EmployeeController) GetEmployees(c *gin.Context)  {
	employees, _ := service.NewEmployeeService().GetEmployees()
	c.JSON(http.StatusOK,&employees)
}

func (e *EmployeeController) GetEmployeeByCode(c *gin.Context) {
	code := c.Param("code")
	employee, err := service.NewEmployeeService().GetEmployeeByCode(code)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, &employee)
	return
}
