package database

import (
	"finalTaskWan/config"
	"finalTaskWan/models"
	"fmt"
)

// SaveData Save the car data in the database
func SaveData(cars []models.Car) error {
	for i := 1; i <= 100; i++ {
		//database query to store the generated into the database
		err := config.DB.Model(&models.Car{}).Select("name", " year", "selling_price", "km_driven", "fuel", "seller_type", "transmission", "owner").Create(&cars).Error
		return err
	}
	fmt.Println(cars)
	return nil
}
