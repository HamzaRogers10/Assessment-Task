package controllers

import (
	"Assessment-Task-Wancloud-Inc/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.User{}
	u.Name = input.Name
	u.Email = input.Email
	u.Password = input.Password
	var err error
	err = models.DB.Model(&models.User{}).Select("name", " email", "password").Create(&u).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}
