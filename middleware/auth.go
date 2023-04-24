package middleware

import (
	"Caro_Game/cache"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func AuthMiddleware(redis *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		result := cache.GetTokenRedis(token, redis)
		if result == "Error" {
			c.JSON(401, gin.H{
				"Message": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthMiddlewareRole(redis *redis.Client) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		result := cache.GetRoleUserRedis(token, redis)
		if result == "1" {
			context.Next()
		} else {
			context.JSON(401, gin.H{
				"Message": "Role is invalid",
			})
			context.Abort()
			return
		}
	}
}
func AuthMiddlewareRoleAdmin(redis *redis.Client) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		result := cache.GetRoleUserRedis(token, redis)
		if result == "2" {
			context.Next()
		} else {
			context.JSON(401, gin.H{
				"Message": "Role ADMIN is invalid",
				"Data":    result,
			})
			context.Abort()
			return
		}
	}
}
