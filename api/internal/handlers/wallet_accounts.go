package handlers

import (
	"arf/currency-conversion/api/app"
	"arf/currency-conversion/api/internal/repository"
	token "arf/currency-conversion/api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type UserWalletAccount struct {
	Balance decimal.Decimal
	Code    string
}

func UserWalletAccounts(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rep := repository.NewR(app.GetDbConn())
	data, err := rep.WalletAccountByUser(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}
