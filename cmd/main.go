package main

import (
	"Assessment-Task-Wancloud-Inc/controllers"
	"Assessment-Task-Wancloud-Inc/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectToDb()

	r := gin.Default()
	userGroup := r.Group("/api")

	userGroup.POST("/register", controllers.Register)
	userGroup.POST("/login", controllers.Login)

	r.Run(":8080")
}
