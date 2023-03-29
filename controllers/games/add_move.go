package games

import (
	"Caro_Game/lib"
	"Caro_Game/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func AddMove(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idGame, err1 := strconv.Atoi(ctx.Param("id_game"))
		if err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "falied id game",
			})
			return
		}
		idPlayer, err2 := strconv.Atoi(ctx.Param("id_player"))
		if err2 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "failed id player",
			})
			return
		}
		var move models.Moves
		err := ctx.BindJSON(&move)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": lib.INVALID_REQUEST,
			})
			return
		}
		move.GameId = idGame
		move.PlayerId = idPlayer

		if err := db.Create(&move).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Add Move is failed",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Add move success",
			"data":    move,
		})
	}
}
