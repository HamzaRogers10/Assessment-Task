package models

import "github.com/jinzhu/gorm"

type Car struct {
	gorm.Model
	Name         string `gorm:"size:255;not null;" json:"name"`
	Year         string `gorm:"size:255;not null;" json:"year"`
	SellingPrice string `gorm:"size:255;not null;" json:"selling_price"`
	KmDriven     string `gorm:"size:255;not null;" json:"km_driven"`
	Fuel         string `gorm:"size:255;not null;" json:"fuel"`
	SellerType   string `gorm:"size:255;not null;" json:"seller_type"`
	Transmission string `gorm:"size:255;not null;" json:"transmission"`
	Owner        string `gorm:"size:255;not null;" json:"owner"`
}
