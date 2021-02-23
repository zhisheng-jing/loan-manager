package global

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB, _ = InitDB()
//连接到数据库
func InitDB() (*gorm.DB, *gorm.DB) {

	DB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	return DB, nil
}