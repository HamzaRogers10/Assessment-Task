package database

import (
	"errors"
	"finalTaskWan/config"
	"finalTaskWan/models"
)

func GetUserByID(uid uint) (models.User, error) {
	var u models.User
	if err := config.DB.First(&u, uid).Error; err != nil {
		return u, errors.New("user not found")
	}
	return u, nil
}
