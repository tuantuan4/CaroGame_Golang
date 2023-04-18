package games

import (
	"Caro_Game/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetHistoryUser(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var result models.User
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id not type",
			})
			return
		}
		if err := db.Where("id = ?", id).First(&result).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id not found",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"username": result.Username,
			"win":      result.Win,
			"lose":     result.Lose,
			"draw":     result.Draw,
		})

	}
}
