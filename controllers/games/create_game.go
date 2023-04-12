package games

import (
	"Caro_Game/lib"
	"Caro_Game/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateGame(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data models.Games
		err := ctx.BindJSON(&data)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": lib.INVALID_REQUEST,
			})
			return
		}

		if err1 := db.Create(&data).Error; err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "failed",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data":    data.ID,
			"message": "create game is success",
		})

	}
}
