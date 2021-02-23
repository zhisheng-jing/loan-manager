package router

import (
	"github.com/gin-gonic/gin"
	"zhisheng-jing/loan-manager/controller"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/employees/getbyid/:id",controller.NewEmployeeController().GetEmployeeByID)
	router.GET("/employees/getbycode/:code",controller.NewEmployeeController().GetEmployeeByCode)
	router.GET("/employees",controller.NewEmployeeController().GetEmployees)
	return router
}
