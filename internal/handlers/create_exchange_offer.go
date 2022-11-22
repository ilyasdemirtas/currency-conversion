package handlers

import (
	"arf/currency-conversion/app"
	"arf/currency-conversion/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExchangeOfferInput struct {
	BaseCurrency    string `json:"base_currency" binding:"required"`
	CounterCurrency string `json:"counter_currency" binding:"required"`
}

func CreateExchangeOffer(c *gin.Context) {

	var input ExchangeOfferInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := app.GetDbConn()
	exchangeRateR := repository.NewExchangeRate(db)
	rate, err := exchangeRateR.GetExchangeRateByCurrencyNames(input.BaseCurrency, input.CounterCurrency)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": rate})

}
