package users

import (
	"Caro_Game/cache"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
)

func Logout(redis *redis.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id_user"))
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err,
			})
			return
		}
		tokenString := cache.GetTokenRedis(id, redis)
		if tokenString == "Error" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Authorization header is required",
			})
			return
		}
		cache.DeleteTokenRedis(id, redis)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Logged out successfully",
		})
		//userID, _ := ctx.Get("UserID")

		//var token models.Token
		//result := db.Find(&token, "user_id = ? AND token = ?", userID, tokenString)
		//db.Delete(&result)
		//afterDeleteStr := ctx.GetHeader("Authorization")
		//if afterDeleteStr == "" {
		//	ctx.JSON(http.StatusOK, gin.H{
		//		"message": "Logged out successfully",
		//	})
		//} else {
		//	ctx.JSON(http.StatusOK, gin.H{
		//		"message": "Logged out failed",
		//	})
		//}

	}
}
