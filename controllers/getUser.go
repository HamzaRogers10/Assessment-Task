package controllers

import (
	"finalTaskWan/jwtAuthentication"
	"finalTaskWan/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CurrentUser(c *gin.Context) {

	userId, err := jwtAuthentication.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "Data": u})
}
