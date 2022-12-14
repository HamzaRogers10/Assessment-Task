package main

import (
	"finalTaskWan/config"
	"finalTaskWan/controllers"
	"finalTaskWan/database"
	"finalTaskWan/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	config.ConnectToDb()
	go backgroundTask()

	r := gin.Default()

	userGroup := r.Group("/api")
	//register our login information, so it can be verified
	userGroup.POST("/register", controllers.Register)
	//This route will be used to authenticate the user by providing  email and password
	//then generate and gives JSON Web Token in return
	userGroup.POST("/login", controllers.Login)
	//this will be the route for our protected endpoints
	userGroup.GET("/user", controllers.CurrentUser)
	userGroup.GET("/cars", database.GetCars)
	userGroup.GET("/filterCars", database.FilterCars)

	userGroup.Use(middleware.JwtAuthMiddleware())

	r.Run(":3000")

}

// Creating a background task function
func backgroundTask() {

	for {

		response := controllers.GetData()
		err := database.SaveData(response)
		fmt.Println(err)
		time.Sleep(time.Hour)
	}
}
