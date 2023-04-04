package middleware

//
//import (
//	"github.com/dgrijalva/jwt-go"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func AuthMiddleware() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		tokenString := context.GetHeader("Authorization")
//		if tokenString == "" {
//			if tokenString == "" {
//				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
//				return
//			}
//
//			// Giải mã JWT và kiểm tra xác thực
//			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//				return []byte("my-secret-key"), nil
//			})
//			if err != nil || !token.Valid {
//				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization token"})
//				return
//			}
//
//			// Lưu thông tin người dùng vào biến Context để sử dụng ở các middleware khác
//			context.Set("user_id", token.Claims.(jwt.MapClaims)["user_id"].(string))
//
//			// Chuyển tiếp đến middleware tiếp theo
//			context.Next()
//		}
//	}
//}
//
//// Middleware để kiểm tra quyền truy cập
//func RoleMiddleware(roles ...string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// Lấy thông tin người dùng từ biến Context
//		userId := c.GetString("user_id")
//
//		// Kiểm tra xem người dùng có quyền truy cập vào API không
//		if !checkUserRole(userId, roles) {
//			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient role"})
//			return
//		}
//
//		// Chuyển tiếp đến middleware tiếp theo
//		c.Next()
//	}
//}
//
