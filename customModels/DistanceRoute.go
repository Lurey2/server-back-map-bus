package custommodels

import "rutasMap/v2/models"

type DistanceRoute struct {
	Distance        float64
	DistanceInitial float64
	DistanceDestiny float64
	DistanceTotal   float64
	PointInitial    models.Point
	PointEnd        models.Point
	Route           models.Route
}

func (d *DistanceRoute) GetRealDistance() float64 {
	return d.DistanceInitial + d.Distance/5 + d.DistanceDestiny
}
