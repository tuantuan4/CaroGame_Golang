package games

import (
	"Caro_Game/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func HistoryRateByUsername(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")

		var user models.User
		if err1 := db.Where("username = ?", username).First(&user).Error; err1 != nil {
			ctx.JSON(400, gin.H{
				"error": "Username not found",
			})
			return
		}
		sum := float64(user.Win + user.Lose + user.Draw)

		if sum == 0 {
			ctx.JSON(200, gin.H{
				"Win":  0,
				"Lose": 0,
				"Draw": 0,
			})
		} else {
			var winRate, loseRate, drawRate float64
			// chua lam tron
			winRate = float64(user.Win) / sum * 100
			loseRate = float64(user.Lose) / sum * 100
			drawRate = float64(user.Draw) / sum * 100
			// da lam tron nhung tra ve dang float64
			//ctx.JSON(200, gin.H{
			//	"Win":  math.Round(winRate*100) / 100,
			//	"Lose": math.Round(loseRate*100) / 100,
			//	"Draw": math.Round(drawRate*100) / 100,
			//	"Sum":  user.Win + user.Lose + user.Draw,
			//})
			//da lam trong, du lieu tra ve dang string strconv.FormatFloat(num, 'f', 2, 64)
			ctx.JSON(200, gin.H{
				"win":  strconv.FormatFloat(winRate, 'f', 2, 64),
				"lose": strconv.FormatFloat(loseRate, 'f', 2, 64),
				"draw": strconv.FormatFloat(drawRate, 'f', 2, 64),
				"sum":  user.Win + user.Lose + user.Draw,
			})
		}

	}
}
