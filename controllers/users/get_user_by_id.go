package users

import (
	"Caro_Game/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetUserById(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id_user"))
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err,
			})
			return
		}
		var user models.User
		if err := db.Where("id = ?", id).First(&user).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id not found",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"data": user.Username,
		})

	}
}
