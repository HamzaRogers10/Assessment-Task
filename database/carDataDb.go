package database

import (
	"finalTaskWan/models"
	"fmt"
)

// SaveData Save the car data in the database
func SaveData(cars []models.Car) error {
	for i := 1; i <= 40; i++ {
		//database query to store the generated into the database
		err := models.DB.Model(&models.Car{}).Select("name", " year", "selling_price", "km_driven", "fuel", "seller_type", "transmission", "owner").Create(&cars).Error
		return err
	}
	fmt.Println(cars)
	return nil
}
