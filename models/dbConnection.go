package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// To save credentials in the database,lets create a Database connection
var DB *gorm.DB
var urlDSN = "root:Hamza@10@tcp(localhost:3306)/carDb?charset=utf8mb4&parseTime=True&loc=Local"
var err error

func ConnectToDb() {
	DB, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	User := User{}
	err = DB.AutoMigrate(&User, &Car{})
	if err != nil {
		return
	}
}
