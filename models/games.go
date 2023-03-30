package models

import "gorm.io/gorm"

type Games struct {
	gorm.Model
	WinnerId int     `json:"winnerId"`
	LoserId  int     `json:"loserId"`
	Moves    []Moves `gorm:"foreignKey:game_id"`
	//Player1  User    `gorm:"ForeignKey:Player1"`
	//Player2  User    `gorm:"ForeignKey:Player2"`
}
