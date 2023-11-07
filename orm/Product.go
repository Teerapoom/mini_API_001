package orm

import (
	"gorm.io/gorm"
)

// โครงสร้าง ตารางใน Databast
type Product struct {
	gorm.Model
	Name   string
	Price  float64
	Detail string
	ImgUrl string
}
