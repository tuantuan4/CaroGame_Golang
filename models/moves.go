package models

import "gorm.io/gorm"

type Moves struct {
	gorm.Model
	XCoordinate int `json:"x_coordinate" validate:"gte=0,lte=2"`
	YCoordinate int `json:"y_coordinate" validate:"gte=0,lte=2"`
	GameId      int `json:"game_id"`
	PlayerId    int `json:"player_id"`
}
