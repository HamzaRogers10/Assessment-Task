package database

import (
	"finalTaskWan/config"
	"finalTaskWan/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FilterCars should be able to query the car dataset based on make model and make year
func FilterCars(c *gin.Context) {
	var cars []models.Car
	name := c.Request.URL.Query().Get("name")
	year := c.Request.URL.Query().Get("year")
	if err := config.DB.Where("name = ? AND year = ?", name, year).Limit(10).Offset(0).Find(&cars).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"dataset": cars})
}
