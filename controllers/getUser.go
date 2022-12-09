package controllers

import (
	"finalTaskWan/jwtAuthentication"
	"finalTaskWan/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CurrentUser we can return the current authenticated user data
func CurrentUser(c *gin.Context) {

	userId, err := jwtAuthentication.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// it authenticated the user by ID , as ID is always unique
	u, err := models.GetUserByID(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//if user is authenticated it will print message success
	c.JSON(http.StatusOK, gin.H{"message": "success", "Data": u})
}
