package main

import (
	inputdata "teerapoom/miniprojset_api_001/Contorller/InputData"
	"teerapoom/miniprojset_api_001/Contorller/output"
	"teerapoom/miniprojset_api_001/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	orm.InitDB()
	r := gin.Default()    //การสร้างอ็อบเจ็กต์ (object) จากคลาส (class) ->  instance
	r.Use(cors.Default()) //การเข้าถึง api ของเรา
	r.POST("/POST_productdata", inputdata.InputProduct)
	r.GET("/GetAll_productdata", output.GetProduct)
	r.GET("/GetbyID_productdata/:id", output.GetProductByID)
	r.PUT("/UpdateProduct/:id", inputdata.UpdateProduct)
	r.DELETE("/DeleteProduct/:id", inputdata.DeleteProduct)
	r.Run("localhost:8080")
}

// json เพราะว่าจะส่งค่าเข้าไปใน api สร้างฟอม
type Product_shop001 struct {
	Name   string  `json:"name" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
	Detail string  `json:"detail" binding:"required"`
	ImgUrl string  `json:"imgUrl" binding:"required"`
}

// โครงสร้าง ตารางใน Databast
type Product struct {
	gorm.Model
	Name   string
	Price  float64
	Detail string
	ImgUrl string
}
