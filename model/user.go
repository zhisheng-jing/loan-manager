package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Employee Employee
	Password string                                //密码
	Username string                                //用户名
	 }
