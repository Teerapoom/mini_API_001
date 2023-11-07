package inputdata

import (
	"net/http"
	"teerapoom/miniprojset_api_001/orm"

	"github.com/gin-gonic/gin"
)

// json เพราะว่าจะส่งค่าเข้าไปใน api สร้างฟอม
type Product_shop001 struct {
	Name   string  `json:"name" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
	Detail string  `json:"detail" binding:"required"`
	ImgUrl string  `json:"imgUrl" binding:"required"`
}

func InputProduct(c *gin.Context) {
	var json Product_shop001 //ตัวแปร json type struct .
	if err := c.ShouldBindJSON(&json); err != nil {
		// เพื่อ bind ข้อมูล JSON ที่ส่งเข้ามาใน request body
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//เช็คซ้ำ
	var ProductExist orm.Product
	orm.Db.Where("Name = ?", json.Name).First(&ProductExist)
	if ProductExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"Status":  "error",
			"Message": "Data Exist",
		})
		return
	}

	//สร้าง
	Product := orm.Product{Name: json.Name, Price: json.Price,
		Detail: json.Detail, ImgUrl: json.ImgUrl}
	orm.Db.Create(&Product)
	if Product.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"Status":      "OK",
			"Message":     "Post Successful",
			"DataProduct": json,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Status":  "error",
			"Message": "Post Failed",
		})
	}
}

func UpdateProduct(c *gin.Context) {
	var product orm.Product
	id := c.Param("id") // รับค่าพารามิเตอร์ "id" จาก URL และเก็บไว้ในตัวแปร id.
	// ตรวจสอบว่ามีสินค้าที่ต้องการอัปเดตหรือไม่
	if err := orm.Db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	// อ่านข้อมูลใหม่จากคำขอ HTTP และอัปเดตสินค้า
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
		return
	}
	// ทำการอัปเดตข้อมูลสินค้าในฐานข้อมูล
	if err := orm.Db.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  "Ok",
		"Product": product})
}

func DeleteProduct(c *gin.Context) {
	var product orm.Product
	id := c.Param("id") // รับค่าพารามิเตอร์ "id" จาก URL และเก็บไว้ในตัวแปร id.

	// ตรวจสอบว่ามีสินค้าที่ต้องการลบหรือไม่
	if err := orm.Db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	// ลบสินค้าออกจากฐานข้อมูล
	if err := orm.Db.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  "Ok",
		"message": "Product deleted successfully"})
}
