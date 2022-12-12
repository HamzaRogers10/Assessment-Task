package controllers

import (
	"finalTaskWan/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register function will register the detail of user
func Register(c *gin.Context) {

	var input models.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.User{}
	u.Name = input.Name
	u.Email = input.Email
	u.Password = input.Password
	var err error
	//DB query to store user data into the database
	err = models.DB.Model(&models.User{}).Select("name", " email", "password").Create(&u).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if the user is register successfully it will print message
	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}
