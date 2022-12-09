package models

import (
	"errors"

	"finalTaskWan/jwtAuthentication"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GetUserByID(uid uint) (User, error) {
	var u User
	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("user not found")
	}
	u.PrepareGive()
	return u, nil
}

// PrepareGive function to remove the hashed password string before returning it for security purposes.
func (u *User) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
func LoginCheck(email string, password string) (string, error) {

	var err error
	u := User{}
	err = DB.Model(&User{}).Where("email = ?", email).First(&u).Error

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
