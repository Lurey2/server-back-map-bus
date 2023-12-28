package services

import (
	"fmt"
	"rutasMap/v2/custom"
	"rutasMap/v2/models"
	"sort"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func GetFindRoute() ([]models.Route, error) {
	var routes []models.Route

	db, err = connection()

	if err != nil {
		fmt.Println("Error")
		return nil, err
	}

	if err := db.Find(&routes).Preload("Routes").Find(&routes).Error; err != nil {
		fmt.Println("Error")
		return nil, err
	}
	for _, element := range routes {

		sort.Sort(custom.NodeSortingIndex(element.Routes))

	}

	return routes, err // aqui se atasca
}

func GetFindActiveRoute() ([]models.Route, error) {
	var routes []models.Route
	var routeTemp []models.Route

	db, err = connection()

	if err != nil {
		fmt.Println("Error")
		return nil, err
	}

	if err := db.Preload("Routes").Find(&routes).Error; err != nil {
		fmt.Println("Error")
		return nil, err
	}
	for _, element := range routes {
		if element.State {
			sort.Sort(custom.NodeSortingIndex(element.Routes))
			routeTemp = append(routeTemp, element)
		}

	}

	return routeTemp, err // aqui se atasca
}

func CreateRoute(route *models.Route) error {

	db, err = connection()

	if err != nil {
		fmt.Println("Error ConnectDb ")
	}

	if err := db.Create(&route).Error; err != nil {
		fmt.Println("Error Create ")
	}
	return err
}

func UpdateRoute(route *models.Route) error {
	db, err = connection()
	if err != nil {
		fmt.Println("Error ConnectDb ")
	}

	go db.Unscoped().Where("route_id = ?", route.ID).Delete(models.Point{})

	if err := db.Model(&route).Updates(route).Error; err != nil {
		fmt.Println("Error Create ")
	}

	db.Model(&route).Select("state").Updates(map[string]interface{}{"state": route.State})

	return err
}

func GetfindByID(route *models.Route, ID int) error {
	db, err = connection()
	if err != nil {
		fmt.Println("Error ConnectDb ")
	}

	if err := db.Preload("Routes").First(&route, ID).Error; err != nil {
		fmt.Println("Error Find ")
	}
	sort.Sort(custom.NodeSortingIndex(route.Routes))
	return err
}

func GetfindByUserID(id uint) []models.Route {
	var routes []models.Route
	db, err = connection()
	if err != nil {
		fmt.Println("Error ConnectDb ")
	}

	if err = db.Where("id_user = ?", id).Find(&routes).Error; err != nil {
		fmt.Println("error ", err)
	}
	return routes
}

func GetRoutePointfindByUserID(id uint) []models.Route {
	var routes []models.Route
	db, _ = connection()

	if err = db.Preload("Routes").Where("id_user = ?", id).Find(&routes).Error; err != nil {
		fmt.Println("error ", err)
	}

	return routes
}

func DeleteRoute(id uint) error {
	db, err = connection()

	if err != nil {
		fmt.Println("Error ConnectDb ")
		return err
	}

	err = db.Delete(&models.Route{}, id).Error
	if err != nil {
		fmt.Println("error ", err)
	}

	return err
}
