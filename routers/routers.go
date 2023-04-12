package routers

import (
	"Caro_Game/controllers/games"
	"Caro_Game/controllers/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DefineRouter(r *gin.Engine, db *gorm.DB) {
	v1 := r.Group("/v1")
	{
		user := v1.Group("users")
		{
			user.POST("/register", users.Register(db))
			user.POST("/login", users.Login(db))
			user.POST("/logout", users.Logout(db))
			user.GET("", users.GetAllUsers(db))
			user.GET("/:id_user", users.GetUserById(db))
		}
		game := v1.Group("/games")
		{
			game.POST("", games.CreateGame(db))                   //done
			game.POST("/AddMove/:id_game", games.AddMove(db))     // done
			game.GET("/:id_game", games.GetGame(db))              // done
			game.GET("/CheckWin/:id_game", games.CheckWin(db))    //done
			game.GET("/GetHistory/:id", games.GetHistoryUser(db)) // done
			game.GET("/rate/:id", games.HistoryRare(db))          // done
			game.GET("/time/:id", games.GetTime(db))              //

		}
	}
}

func InitializeRouters(db *gorm.DB) {
	r := gin.Default()
	DefineRouter(r, db)
	r.Run()
}
