package users

import (
	"Caro_Game/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Logout(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, _ := ctx.Get("user")
		if u, ok := user.(models.User); ok {
			userID := u.ID
			tokenString := ctx.GetHeader("Authorization")
			if tokenString == "" {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "Authorization header is required",
				})
				return
			}
			var token models.Token
			result := db.Where("user_id ? AND token = ?", userID, tokenString).First(&token)
			if result.Error != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid token",
				})
			}
			db.Delete(&token)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Logged out successfully",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
			})
		}

	}
}
