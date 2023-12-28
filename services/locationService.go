package services

import (
	"fmt"
	"rutasMap/v2/conf"
	"rutasMap/v2/models"

	"gorm.io/gorm"
)

func connection() (db *gorm.DB, err error) {
	if db == nil {
		return conf.InitDb()
	}
	return db, nil
}

func GetFindLocation() ([]models.Location, error) {
	var locations []models.Location

	db, err = connection()

	if err != nil {
		fmt.Println("Error")
		return nil, err
	}

	if err := db.Find(&locations).Error; err != nil {
		fmt.Println("Error")
		return nil, err
	}
	
	return locations, err // aqui se atasca
}


func CreateLocation(location *models.Location) error {

	db, err = connection()

	if err != nil {
		fmt.Println("Error ConnectDb ")
	}

	if err := db.Create(&location).Error; err != nil {
		fmt.Println("Error Create ")
	}
	return err
}

func UpdateLocation(location *models.Location) error {
	db, err = connection()
	if err != nil {
		fmt.Println("Error ConnectDb ")
	}
	
	if err := db.Save(&location).Error; err != nil {
		fmt.Println("Error Update ")
	}

	return err
}

func GetfindByIDLocation(location *models.Location, ID int) error {
	db, err = connection()
	if err != nil {
		fmt.Println("Error ConnectDb ")
	}

	if err := db.First(&location, ID).Error; err != nil {
		fmt.Println("Error Find ")
	}
	return err
}
