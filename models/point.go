package models

import "gorm.io/gorm"

type Point struct {
	gorm.Model
	Index     uint
	IndexNext uint
	Distance  float64
	Lat       float64
	Lng       float64
	Show      bool
	RouteId   uint
}
