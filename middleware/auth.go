package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		isLoggedIn := false
		tokenStr, _ := context.Get("Authorization")
		if tokenStr != "" {
			isLoggedIn = true
		}
		if !isLoggedIn {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "please login",
			})
			context.Abort()
			return
		}
		context.Next()
	}
}

//func AuthMiddleware() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		tokenString, _ := context.Get("Authorization")
//		if tokenString == "" {
//			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
//			return
//		}
//		//Giải mã JWT và kiểm tra xác thực
//		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		//	return []byte("my-secret-key"), nil
//		//})
//		//if err != nil || !token.Valid {
//		//	context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization token"})
//		//	return
//		//}
//		// Lưu thông tin người dùng vào biến Context để sử dụng ở các middleware khác
//		//context.Set("user_id", token.Claims.(jwt.MapClaims)["user_id"].(string))
//
//		// Chuyển tiếp đến middleware tiếp theo
//		//context.Next()
//	}
//}
//
//// // Middleware để kiểm tra quyền truy cập
//func RoleMiddleware(roles ...string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// Lấy thông tin người dùng từ biến Context
//		//userId, _ := c.Get("UserID")
//		roleId, _ := c.Get("RoleID")
//		// Kiểm tra xem người dùng có quyền truy cập vào API không
//
//		if roleId == 2 {
//			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient role"})
//			return
//		}
//		// Chuyển tiếp đến middleware tiếp theo
//		c.Next()
//	}
//}
