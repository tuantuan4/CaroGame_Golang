package models

import "gorm.io/gorm"

type Games struct {
	gorm.Model
	PlayerId1 int     `json:"player_id1"`
	PlayerId2 int     `json:"player_id2"`
	WinnerId  int     `json:"winnerId"`
	LoserId   int     `json:"loserId"`
	Moves     []Moves `gorm:"foreignKey:game_id"`
	//Player1  User    `gorm:"ForeignKey:Player1"`
	//Player2  User    `gorm:"ForeignKey:Player2"`
}
