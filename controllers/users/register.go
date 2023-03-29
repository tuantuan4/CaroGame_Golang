package users

import (
	"Caro_Game/lib"
	"Caro_Game/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var user models.User
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": lib.INVALID_REQUEST,
			})
			return
		}
		//valid username
		var count int64
		db.Model(&models.User{}).Where("username = ?", user.Username).Count(&count)
		if count > 0 {
			ctx.JSON(400, gin.H{
				"error": lib.USER_ALREADY_EXISTS,
			})
			return
		}
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": lib.FLAIED_TO_HASH_PASSWORD,
			})
			return
		}
		user.Password = string(passwordHash)
		//them thong tin vao db
		result := db.Create(&user)
		if result.Error != nil {
			ctx.JSON(400, gin.H{
				"error": lib.FLAIED_TO_CREATE_USER,
			})
			return
		}
		ctx.JSON(200, user)
	}
}
