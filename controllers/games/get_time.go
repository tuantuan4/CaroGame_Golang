package games

import (
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
		//var game models.Games
		//if err := db.Where("id = ?", id).First(&game).Error; err != nil {
		//	ctx.JSON(http.StatusBadRequest, gin.H{
		//		"error": "id not found",
		//	})
		//	return
		//}
		var sum int64
		if err1 := db.Table("games").Select("sum(updated_at - created_at)").
			Where("player_id1 = ?", id).Or("player_id2 = ?", id).
			Row().Scan(&sum).Error; err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to render time",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"data": sum,
		})
	}
}
