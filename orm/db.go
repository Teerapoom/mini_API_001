package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitDB() {
	// เชื่อม Data bast
	dsn := "root:1234@tcp(127.0.0.1:3306)/product_miniapi001?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	Db.AutoMigrate(&Product{})
	// เชื่อม Data bast
}
