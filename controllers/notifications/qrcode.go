package notifications

import (
	"Caro_Game/common"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"image/png"
	"net/http"
	"os"
)

var reqBody struct {
	Data string `json:"data" binding:"required"`
}

func GenerateQRCode(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		qrCode, err := qr.Encode(reqBody.Data, qr.L, qr.Auto)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate QR code"})
			return
		}
		filename := "images/" + common.RandomString() + ".png"
		file, err := os.Create(filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
			return
		}
		defer file.Close()
		png.Encode(file, qrCode)
		c.File(filename)
		//var buf bytes.Buffer
		//if err := png.Encode(&buf, qrCode); err != nil {
		//	c.JSON(500, gin.H{"error": "Failed to encode QR code as PNG"})
		//	return
		//}
		//c.Data(200, "image/png", buf.Bytes())
	}
}
