package models

import "gorm.io/gorm"

type RouteShow struct {
	gorm.Model
	Route Route `gorm:"not null"`
	Order uint8 `gorm:"not null"`
}
