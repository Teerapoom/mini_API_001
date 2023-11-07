package output

import (
	"net/http"
	"teerapoom/miniprojset_api_001/orm"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	var Product []orm.Product
	orm.Db.Find(&Product)
	c.JSON(http.StatusOK, gin.H{"Product": Product})
}

func GetProductByID(c *gin.Context) {
	var product orm.Product
	id := c.Param("id") //รับค่าพารามิเตอร์ "id" จาก URL และเก็บไว้ในตัวแปร id.

	// First เป็นคำสั่งการค้นหา
	if err := orm.Db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return //Stop
	}

	c.JSON(http.StatusOK, gin.H{"Product": product})
}
