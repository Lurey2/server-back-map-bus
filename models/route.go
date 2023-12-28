package models

import (
	"gorm.io/gorm"
)

type Route struct {
	gorm.Model
	Name        string
	Color       string
	State       bool `gorm:"column:state"`
	Description string
	IdUser      uint    `gorm:"->;<-:create"`
	Routes      []Point `gorm:"foreignKey:RouteId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (r *Route) FindIndex(index uint) Point {
	var point Point
	for _, pointIterator := range r.Routes {
		if pointIterator.Index == index {
			return pointIterator
		}
	}
	return point
}

func (r *Route) SortPoint() {
	points := [][]Point{}

	for _, v := range r.Routes {
		points = append(points, []Point{v})
	}

}
