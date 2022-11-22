package handlers

import (
	"arf/currency-conversion/app"
	"arf/currency-conversion/internal/repository"
	token "arf/currency-conversion/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {

	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := app.GetDbConn()
	userR := repository.NewUser(db)
	u, err := userR.GetUserByID(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
