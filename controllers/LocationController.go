package controllers

import (
	"net/http"

	"rutasMap/v2/models"
	"rutasMap/v2/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLocation(c *gin.Context) {

	var locations []models.Location
	var err error

	locations, err = services.GetFindLocation()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, locations)
}

func CreateLocation(c *gin.Context) {

	var location models.Location

	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateLocation(&location)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, location)
}

func UpdateLocation(c *gin.Context) {
	var locations models.Location

	if err := c.ShouldBindJSON(&locations); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.UpdateLocation(&locations)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, locations)
}

func FindByIDLocation(c *gin.Context) {
	var locations models.Location

	IDdata, errConvert := strconv.Atoi(c.Param("ID"))
	if errConvert != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errConvert.Error()})
		return
	}
	err := services.GetfindByIDLocation(&locations, IDdata)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, locations)
}
