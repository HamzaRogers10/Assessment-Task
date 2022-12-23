package controllers

import (
	"finalTaskWan/config"
	"finalTaskWan/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCars(c *gin.Context) {
	var cars []models.NewCars
	config.DB.Find(&cars)

	c.JSON(http.StatusOK, gin.H{"data": cars})
}
