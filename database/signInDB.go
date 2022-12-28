package database

import (
	"finalTaskWan/config"
	"finalTaskWan/jwtAuthentication"
	"finalTaskWan/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {

	var err error
	u := models.User{}
	err = config.DB.Model(&models.User{}).Where("email = ?", email).First(&u).Error

	if err != nil {
		log.Println("this is error", err)

		return "", err

	}
	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println("this is error", err)
		return "", err
	}

	token, err := jwtAuthentication.GenerateToken(u.ID)

	if err != nil {
		log.Println("this is error", err)
		return "", err
	}
	return token, nil
}
