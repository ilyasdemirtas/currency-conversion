package handlers

import (
	"arf/currency-conversion/api/app"
	"arf/currency-conversion/api/internal/models"
	"arf/currency-conversion/api/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
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

	// Get ExchangeRate by given currencies
	rate, err := GetExchangeRateByCurrencyNames(input.BaseCurrency, input.CounterCurrency)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate data with ExchangeRate data
	exchangeRateOffer := models.ExchangeRateOffer{
		BaseCurrencyId:    rate.BaseCurrencyId,
		CounterCurrencyId: rate.CounterCurrencyId,
		Price:             CalculateMarkup(rate.Price, rate.MarkupRate),
		MarkupRate:        rate.MarkupRate,
	}

	// Create new offer
	rep := repository.NewR(app.GetDbConn())
	result, err := rep.CreateExchangeOffer(exchangeRateOffer)

	type ExchangeRateOffer struct {
		Price      decimal.Decimal
		MarkupRate decimal.Decimal
	}

	offer := ExchangeRateOffer{
		Price:      result.Price,
		MarkupRate: result.MarkupRate,
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": offer})

}

type ExchangeRateOffer struct {
	BaseCurrencyId    uint32
	CounterCurrencyId uint32
	Price             decimal.Decimal
	MarkupRate        decimal.Decimal
}

func GetExchangeRateOfferByCurrencyNames(baseCurrency, counterCurrency string) (models.ExchangeRateOffer, error) {

	rep := repository.NewR(app.GetDbConn())
	data, err := rep.GetExchangeRateOfferByCurrencyNames(baseCurrency, counterCurrency)

	if err != nil {
		return data, err
	}

	return data, nil

}

func CalculateMarkup(price, markup decimal.Decimal) decimal.Decimal {
	return price.Add(price.Div(decimal.NewFromFloat(100)).Mul(markup))
}
