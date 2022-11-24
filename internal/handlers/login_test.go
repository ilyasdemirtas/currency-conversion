package handlers

import (
	"arf/currency-conversion/internal/database/connection"
	"arf/currency-conversion/internal/models"
	"arf/currency-conversion/internal/repository"
	"testing"
	"unicode/utf8"

	"github.com/joho/godotenv"
)

func TestLogin(t *testing.T) {

	u := models.User{
		Email:    "ilyasdemirtas@hotmail.com.tr",
		Password: "123qwe",
	}

	if err := godotenv.Load("../../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	db := connection.Init()
	repository := repository.NewR(db)
	token, err := repository.LoginCheck(u.Email, u.Password)

	if err != nil {
		t.Errorf("Login error = %v;", err.Error())
	}

	if !utf8.ValidString(token) {
		t.Errorf("Invalid token = %q;", token)
	}
}
