package repository

import (
	"arf/currency-conversion/internal/models"
	token "arf/currency-conversion/internal/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (r R) LoginCheck(email string, password string) (string, error) {

	var err error

	u := models.User{}

	err = r.db.Model(models.User{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	generateToken, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return generateToken, nil

}

func (r R) GetUserByID(uid uint32) (models.User, error) {

	var data models.User

	if err := r.db.First(&data, uid).Error; err != nil {
		return data, errors.New("user not found")
	}

	data.Password = ""

	return data, nil

}
