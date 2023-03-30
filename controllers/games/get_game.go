package games

import (
	"Caro_Game/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetGame(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id_game"))
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err,
			})
			return
		}
		var game models.Games
		if err := db.Where("id = ?", id).First(&game).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id not found",
			})
			return
		}
		var moves []models.Moves

		if err := db.Where("game_id", id).Find(&moves).Error; err != nil {
			ctx.JSON(400, gin.H{
				"error": "failed to load moves",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"game":  game,
			"moves": moves,
		})

	}
}
