package routers

import (
	"Caro_Game/controllers/games"
	"Caro_Game/controllers/users"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func DefineRouter(r *gin.Engine, db *gorm.DB, redis *redis.Client) {
	v1 := r.Group("/v1")
	{
		user := v1.Group("users")
		{
			user.POST("/register", users.Register(db))
			user.POST("/login", users.Login(db, redis))
			user.POST("/logout/:id_user", users.Logout(redis))
			user.GET("/:id_user", users.GetUserById(db))
		}
		game := v1.Group("/games")
		{
			game.POST("", games.CreateGame(db))                             // done
			game.POST("/AddMove/:id_game", games.AddMove(db))               // done
			game.GET("/:id_game", games.GetGame(db))                        // done
			game.GET("/CheckWin/:id_game", games.CheckWin(db))              // done
			game.GET("/GetHistory/:id", games.GetHistoryUser(db))           // done
			game.GET("/rate/:id", games.HistoryRate(db))                    // done
			game.GET("/time/:id", games.GetTime(db))                        // done
			game.GET("/history/:username", games.HistoryRateByUsername(db)) // done
			game.GET("/move/:id_game", games.GetMoveByGame(db))

		}

	}
}

func InitializeRouters(db *gorm.DB, redis *redis.Client) {
	r := gin.Default()
	DefineRouter(r, db, redis)
	r.Run()
}
