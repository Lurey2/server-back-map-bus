package util

import (
	"math"
	"rutasMap/v2/models"
)

func DistanceDestiny(p1 models.Point, p2 models.Point, route models.Route) float64 {
	currentPointIndex := p1.Index
	totalDistance := 0.0
	for {
		point := route.FindIndex(currentPointIndex)
		totalDistance = totalDistance + point.Distance
		if point.IndexNext == p2.Index {
			return (totalDistance)
		}
		currentPointIndex = point.IndexNext

	}

}

func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func DistanceTwoPoints(p1 models.Point, p2 models.Point) float64 {
	earthRadiusInMeters := 6371.0 // Radio de la tierra en metros

	lat1Rad := degToRad(p1.Lat)
	lat2Rad := degToRad(p2.Lat)
	deltaLat := degToRad(p2.Lat - p1.Lat)
	deltaLon := degToRad(p2.Lng - p1.Lng)

	a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(deltaLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusInMeters * c * 1000 * 4
}
