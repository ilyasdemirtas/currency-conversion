package handlers

import (
	"arf/currency-conversion/api/internal/database/connection"
	"arf/currency-conversion/api/internal/models"
	"arf/currency-conversion/api/internal/repository"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
)

func TestGetExchangeRateOfferByCurrencyNames(t *testing.T) {

	if err := godotenv.Load("../../../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	exampleOffer := models.ExchangeRateOffer{
		Price:             decimal.NewFromFloat(18.5383),
		MarkupRate:        decimal.NewFromFloat(0.4),
		BaseCurrencyId:    2,
		CounterCurrencyId: 1,
	}

	db := connection.Init()
	rep := repository.NewR(db)

	offer, err := rep.GetExchangeRateByCurrencyNames("EUR", "TRY")

	if err != nil {
		t.Errorf("Offer not found = %d;", err)
	}

	assert.Equal(t, exampleOffer, offer)

	offer2, err := rep.GetExchangeRateOfferByCurrencyNames("USD", "USD")

	if err != nil {
		t.Errorf("Offer not found = %v;", err.Error())
	}

	assert.NotEqual(t, exampleOffer, offer2)

}
