package handlers

import (
	"arf/currency-conversion/app"
	"arf/currency-conversion/internal/models"
	"arf/currency-conversion/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Email = input.Email
	u.Password = input.Password

	db := app.GetDbConn()
	userR := repository.NewUser(db)
	token, err := userR.LoginCheck(u.Email, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
