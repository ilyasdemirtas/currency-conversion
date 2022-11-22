package handlers

import (
	"arf/currency-conversion/app"
	"arf/currency-conversion/internal/repository"
	token "arf/currency-conversion/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type UserWalletAccount struct {
	Balance decimal.Decimal
	Code    string
}

func UserWalletAccounts(c *gin.Context) {
	db := app.GetDbConn()
	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	walletAccountR := repository.NewWalletAccount(db)
	data, err := walletAccountR.WalletAccountByUser(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}
