package controllers

import (
	"finalTaskWan/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//  it will receive a email and password
func Login(c *gin.Context) {

	// it matches the credential in our database,
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Email = input.Email
	u.Password = input.Password
	//if it does return a token, if it doesn't return an error response
	token, err := models.LoginCheck(u.Email, u.Password)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
