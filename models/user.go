package models

import (
	"github.com/jinzhu/gorm"
)

// User structs
type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null;" json:"name"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
}
