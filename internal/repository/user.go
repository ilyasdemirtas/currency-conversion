package repository

import (
	"arf/currency-conversion/internal/models"
	token "arf/currency-conversion/internal/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserR struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserR {
	return UserR{
		db: db,
	}
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (r UserR) LoginCheck(email string, password string) (string, error) {

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

func (r UserR) GetUserByID(uid uint32) (models.User, error) {

	var u models.User

	if err := r.db.First(&u, uid).Error; err != nil {
		return u, errors.New("user not found")
	}

	u.Password = ""

	return u, nil

}
