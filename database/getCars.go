package database

import (
	"finalTaskWan/config"
	"finalTaskWan/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCars(c *gin.Context) {
	var cars []models.Car
	config.DB.Limit(10).Offset(0).Find(&cars)

	c.JSON(http.StatusOK, gin.H{"data": cars})
}
