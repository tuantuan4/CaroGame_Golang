package games

import (
	"Caro_Game/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func HistoryRare(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err,
			})
			return
		}
		var user models.User
		if err1 := db.Where("id = ?", id).First(&user).Error; err1 != nil {
			ctx.JSON(400, gin.H{
				"error": "ID not found",
			})
			return
		}
		sum := user.Win + user.Lose + user.Draw

		if sum == 0 {
			ctx.JSON(200, gin.H{
				"Win":  0,
				"Lose": 0,
				"Draw": 0,
			})
		} else {
			var winRate, loseRate, drawRate float64
			winRate = float64((user.Win / sum) * 100)
			loseRate = float64((user.Lose / sum) * 100)
			drawRate = float64((user.Draw / sum) * 100)
			ctx.JSON(200, gin.H{
				"Win":  winRate,
				"Lose": loseRate,
				"Draw": drawRate,
			})
		}

	}
}
