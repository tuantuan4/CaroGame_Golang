package config

import (
	"Caro_Game/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

//const (
//	DATABASE_CONNECTION string = "root:tuan@tcp(127.0.0.1:3308)/CaroGame?charset=utf8mb4&parseTime=True&loc=Local"
//)

func ConnectionDatabase() (db *gorm.DB) {
	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(models.User{}, models.Token{}, models.Games{}, models.Moves{})
	return db
}
