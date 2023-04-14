package games

import (
	"Caro_Game/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetMoveByGame(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id_game"))
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err,
			})
			return
		}

		var move []models.Moves
		if err := db.Where("game_id = ?", id).Find(&move).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id_game not found",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"data": move,
		})
	}
}
