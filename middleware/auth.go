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
	//	// Lấy mã thông báo từ header của yêu cầu
	//	authHeader := c.GetHeader("Authorization")
	//	if authHeader == "" {
	//		// Trả về mã trạng thái không cho phép truy cập nếu không có mã thông báo
	//		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
	//		c.Abort()
	//		return
	//	}
	//
	//	// Tách mã thông báo từ header
	//	tokenString := authHeader[len("Bearer "):]
	//
	//	// Giải mã mã thông báo và kiểm tra tính hợp lệ
	//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//		// Kiểm tra loại thuật toán và trả về JWT key
	//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	//		}
	//		return jwtKey, nil
	//	})
	//
	//	if err != nil {
	//		// Trả về mã trạng thái không cho phép truy cập nếu không thể giải mã mã thông báo
	//		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
	//		c.Abort()
	//		return
	//	}
	//
	//	// Lưu thông tin người dùng được lấy từ mã thông báo vào context để sử dụng cho các yêu cầu tiếp theo
	//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	//		userID := claims["id"].(string)
	//		c.Set("userID", userID)
	//		c.Next()
	//	} else {
	//		// Trả về mã trạng thái không cho phép truy cập nếu mã thông báo không hợp lệ
	//		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
	//		c.Abort()
	//		return
	//	}
	//}
}
