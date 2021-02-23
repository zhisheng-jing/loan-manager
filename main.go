package main

import (
	"zhisheng-jing/loan-manager/global"
	"zhisheng-jing/loan-manager/model"
	"zhisheng-jing/loan-manager/router"
)

func init() {
	global.DB.AutoMigrate(&model.Employee{})
}


func main() {

	r := router.Router()
	r.Run(":5000")
}
