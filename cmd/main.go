package main

import (
	"finalTaskWan/controllers"
	"finalTaskWan/middleware"
	"finalTaskWan/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	models.ConnectToDb()
	go backgroundTask()

	r := gin.Default()
	userGroup := r.Group("/api")

	userGroup.POST("/register", controllers.Register)
	userGroup.POST("/login", controllers.Login)
	//userGroup.GET("/filter", controllers.)

	protected := r.Group("/api/admin")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":3000")
}

func backgroundTask() {

	for {

		rsp := controllers.GetData()
		err := controllers.SaveData(rsp)
		fmt.Println(err)
		time.Sleep(time.Hour)
	}
}
