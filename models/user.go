package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Token    Token  `gorm:"foreignKey:UserID"`
	Win      int    `json:"win"`
	Lose     int    `json:"lose"`
	Draw     int    `json:"draw"`

	//Games []Games
	Moves  []Moves `gorm:"foreignKey:PlayerId"`
	RoleID uint    `gorm:"ForeignKey:RoleID"`
	Role   Role
}
