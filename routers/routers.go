package routers

import (
	"Caro_Game/controllers/games"
	"Caro_Game/controllers/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeRouters(db *gorm.DB) {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		user := v1.Group("users")
		{
			user.POST("/register", users.Register(db))
			user.POST("/login", users.Login(db))
			user.POST("/logout", users.Logout(db))
			user.GET("", users.GetAllUsers(db))
		}
		game := v1.Group("/games")
		{
			game.POST("/CreateGame", games.CreateGame(db))               //done
			game.POST("/AddMove/:id_game/:id_player", games.AddMove(db)) // done
			game.GET("/GetGame/:id_game", games.GetGame(db))             // done
			game.GET("/CheckWin/:id_game", games.CheckWin(db))           //done
			game.GET("/GetHistory/:id", games.GetHistoryUser(db))        // done
			game.GET("/rate/:id", games.HistoryRare(db))                 // done

		}
	}
	r.Run()
}
