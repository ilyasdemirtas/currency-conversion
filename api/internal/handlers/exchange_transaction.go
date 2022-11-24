package handlers

import (
	"arf/currency-conversion/api/app"
	"arf/currency-conversion/api/internal/models"
	"arf/currency-conversion/api/internal/repository"
	token "arf/currency-conversion/api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExchangeTransactionInput struct {
	BaseCurrency    string `json:"base_currency" binding:"required"`
	CounterCurrency string `json:"counter_currency" binding:"required"`
}

func CreateExchangeTransaction(c *gin.Context) {
	var input ExchangeTransactionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offer, err := GetExchangeRateOfferByCurrencyNames(input.BaseCurrency, input.CounterCurrency)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exchangeTransaction := models.ExchangeTransaction{
		BaseCurrencyId:    offer.BaseCurrencyId,
		CounterCurrencyId: offer.CounterCurrencyId,
		Price:             offer.Price,
		MarkupRate:        offer.MarkupRate,
		UserId:            userId,
	}

	rep := repository.NewR(app.GetDbConn())
	data, err := rep.CreateExchangeTransaction(exchangeTransaction)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})

}
