package main

import (
	"Caro_Game/config"
	"Caro_Game/routers"
)

func main() {
	db := config.ConnectionDatabase()
	routers.InitializeRouters(db)
}
