package main

import (
	"Caro_Game/cache"
	"Caro_Game/config"
	"Caro_Game/routers"
)

func main() {
	db := config.ConnectionDatabase()
	rdb := cache.NewClientRedis()
	routers.InitializeRouters(db, rdb)

}
