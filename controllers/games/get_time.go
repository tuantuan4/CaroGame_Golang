package games

import (
	"Caro_Game/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetTime(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
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
		var hour int
		if err1 := db.Select("id, created_at, updated_at, HOUR(TIMEDIFF(updated_at, created_at)) as hours_diff").Scan(&hour).Error; err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "filed to convert",
			})
		}
		ctx.JSON(200, gin.H{
			"data": hour,
		})
	}
}
