package games

import (
	"Caro_Game/common"
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

		var sum float64

		if err := db.Raw("SELECT SUM(updated_at - created_at) "+
			"FROM games WHERE player_id1 = ? OR player_id2 = ?", id, id).
			Scan(&sum).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to render time",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"data": common.ConvertSecondsToHMS(sum),
		})
	}
}
