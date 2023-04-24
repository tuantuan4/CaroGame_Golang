package users

import (
	"Caro_Game/cache"
	"Caro_Game/lib"
	"Caro_Game/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Claims struct {
	UserID uint `json:"id"`
	jwt.StandardClaims
}

func generateToken(user models.User) (string, error) {
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // hết hạn sau 24 giờ
		},
	}
	// Tạo token từ claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Ký token với key secret
	return token.SignedString([]byte("mysecretkey")) // đổi lại key secret thực tế của bạn
}

func Login(db *gorm.DB, redis *redis.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var user models.User
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": lib.INVALID_REQUEST,
			})
			return
		}
		// kiem tra thong tin user
		var existingUser models.User
		result := db.Where("username = ?", user.Username).First(&existingUser)
		if result.Error != nil {
			ctx.JSON(400, gin.H{"error": lib.ACCOUNT_IS_INCORRECT})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
			ctx.JSON(400, gin.H{"error": lib.ACCOUNT_IS_INCORRECT})
			return
		}
		tokenString, err := generateToken(existingUser)
		if err != nil {
			ctx.JSON(500, gin.H{"error": lib.TOKEN_ERROR})
			return
		}
		cache.AddToken(int(existingUser.ID), tokenString, redis)
		//cache.AddUserRedis(tokenString, existingUser, redis)
		tokenRecord := models.Token{
			Token:  tokenString,
			UserID: existingUser.ID,
		}
		if err := db.Create(&tokenRecord).Error; err != nil {
			ctx.JSON(500, gin.H{"error": lib.TOKEN_ERROR})
			return
		}
		//tra ve ma token

		ctx.JSON(200, gin.H{
			"token":   tokenString,
			"ID":      existingUser.ID,
			"message": lib.LOGIN_SUCCESS,
		})
	}
}
