package users

import (
	"Caro_Game/cache"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

func Logout(redis *redis.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		result := cache.DeleteTokenRedis(tokenString, redis)
		if result == "error" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Authorization header is required",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Logged out successfully",
		})

	}
}
