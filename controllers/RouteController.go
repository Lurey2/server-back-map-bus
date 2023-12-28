package controllers

import (
	"net/http"
	"rutasMap/v2/conf"
	"rutasMap/v2/custom"
	custommodels "rutasMap/v2/customModels"
	"rutasMap/v2/models"
	"rutasMap/v2/services"
	"rutasMap/v2/util"
	"sort"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

func GetRoute(c *gin.Context) {
	wg := sync.WaitGroup{}
	var routes []models.Route
	var err error

	routes, err = services.GetFindRoute()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, element := range routes {
		wg.Add(1)

		sorter := func(element models.Route) {
			sort.Sort(custom.NodeSortingIndex(element.Routes))
			wg.Done()
		}

		go sorter(element)

	}
	wg.Wait()

	c.JSON(http.StatusOK, routes)
}

func CreateRoute(c *gin.Context) {

	var route models.Route
	var user models.User

	if err := c.ShouldBindJSON(&route); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cook, _ := c.Cookie("SESSIONID")
	data, _ := conf.DecodeToken(cook)
	user.ConvertMapStruct(data.Data)

	route.IdUser = user.ID

	err := services.CreateRoute(&route)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, route)
}

func UpdateRoute(c *gin.Context) {
	var routes models.Route
	var user models.User

	cook, _ := c.Cookie("SESSIONID")
	data, _ := conf.DecodeToken(cook)
	user.ConvertMapStruct(data.Data)

	if err := c.ShouldBindJSON(&routes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.ID != routes.IdUser {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not authorization"})
		return
	}

	routes.SortPoint()

	err := services.UpdateRoute(&routes)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, routes)
}

func FindByID(c *gin.Context) {
	var routes models.Route

	IDdata, errConvert := strconv.Atoi(c.Param("ID"))
	if errConvert != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errConvert.Error()})
		return
	}
	err := services.GetfindByID(&routes, IDdata)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, routes)
}

func FindRouteByUserID(c *gin.Context) {
	var routes []models.Route
	var user models.User

	cook, _ := c.Cookie("SESSIONID")
	data, _ := conf.DecodeToken(cook)
	user.ConvertMapStruct(data.Data)

	routes = services.GetfindByUserID(user.ID)

	c.JSON(http.StatusOK, conf.Response{
		Status: conf.Succes,
		Data:   routes,
	})
}

func FindRoutePointByUser(c *gin.Context) {
	var routes []models.Route
	var user models.User

	cook, _ := c.Cookie("SESSIONID")
	data, _ := conf.DecodeToken(cook)
	user.ConvertMapStruct(data.Data)

	routes = services.GetRoutePointfindByUserID(user.ID)
	for _, element := range routes {

		sort.Sort(custom.NodeSortingIndex(element.Routes))

	}
	c.JSON(http.StatusOK, conf.Response{
		Status: conf.Succes,
		Data:   routes,
	})
}

func FindBYNearbyRoute(c *gin.Context) {
	var contextPoint util.OrigenDestinyRoute

	if err := c.ShouldBindJSON(&contextPoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rutas, err := services.GetFindActiveRoute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var distanceRoute []custommodels.DistanceRoute
	for _, route := range rutas {
		
		if len(route.Routes) > 0 {
			pointOrigin := models.Point{Lat: contextPoint.CoordInitialLat, Lng: contextPoint.CoordInitialLng}
			pointDestiny := models.Point{Lat: contextPoint.CoordEndLat, Lng: contextPoint.CoordEndLng}
			var pointInitial models.Point
			var pointEnd models.Point
			distanceInitial := 0.0
			distanceEnd := 0.0
			for _, pointIndex := range route.Routes {

				if distanceInitial == 0 || distanceInitial > (util.DistanceTwoPoints(pointOrigin, pointIndex)) {
					distanceInitial = util.DistanceTwoPoints(pointOrigin, pointIndex)
					pointInitial = pointIndex

				}

				if distanceEnd == 0 || distanceEnd > (util.DistanceTwoPoints(pointDestiny, pointIndex)) {
					distanceEnd = util.DistanceTwoPoints(pointDestiny, pointIndex)
					pointEnd = pointIndex
				}
			}
			
			totalDistace := util.DistanceDestiny(pointInitial, pointEnd, route)
			distanceRoute = append(distanceRoute, custommodels.DistanceRoute{Route: route, Distance: (totalDistace), DistanceInitial: distanceInitial, DistanceDestiny: distanceEnd, DistanceTotal: distanceInitial + totalDistace/4.5 + distanceEnd, PointInitial: pointInitial, PointEnd: pointEnd})

		}

	}
	sort.SliceStable(distanceRoute, func(i, j int) bool {
		return distanceRoute[i].GetRealDistance() < distanceRoute[j].GetRealDistance()
	})

	c.JSON(http.StatusOK, distanceRoute)
}

func FindRouteNearbyLatlng(c *gin.Context) {

	var contextPoint util.OrigenDestinyRoute

	if err := c.ShouldBindJSON(&contextPoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rutas, err := services.GetFindActiveRoute()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pointRef := models.Point{Lat: contextPoint.CoordInitialLat, Lng: contextPoint.CoordInitialLng}

	listOrder := []util.PointDistance{}
	for _, route := range rutas {
		var distance float64 = 0
		var point *models.Point
		for _, points := range route.Routes {
			distancePoint := util.DistanceTwoPoints(pointRef, points)

			if distance > distancePoint || distance == 0 {
				point = &points
				distance = distancePoint
			}
		}
		if distance == 0 {
			continue
		}
		listOrder = append(listOrder, util.PointDistance{Route: route, Point: *point, Distance: distance})
	}
	sort.SliceStable(listOrder, func(i, j int) bool {
		return listOrder[i].Distance < listOrder[j].Distance
	})
	c.JSON(http.StatusOK, listOrder)
}

func GetActiveRoute(c *gin.Context) {
	var routes []models.Route
	var err error
	wg := sync.WaitGroup{}

	routes, err = services.GetFindActiveRoute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, element := range routes {

		wg.Add(1)

		sorter := func(element models.Route) {
			sort.Sort(custom.NodeSortingIndex(element.Routes))
			wg.Done()
		}

		go sorter(element)

	}
	wg.Wait()
	c.JSON(http.StatusOK, routes)
}

func DeleteRoute(c *gin.Context) {

	var id int
	var err error

	id, err = strconv.Atoi(c.Param("ID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	err = services.DeleteRoute(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete route",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
	return
}
