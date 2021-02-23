package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB, _ = InitDB()
//连接到数据库
func InitDB() (*gorm.DB, *gorm.DB) {
	dsn := "root:000000@tcp(127.0.0.1:3306)/loan-manager?charset=utf8&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	return DB, nil
}