package models

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Lat     float64	`gorm:"not null"`
	Lng     float64	`gorm:"not null"`
	Name	string	`gorm:"not null;size:100"`
	Icon	string  `gorm:"not null"`
	Color 	string	`gorm:"not null"`

}