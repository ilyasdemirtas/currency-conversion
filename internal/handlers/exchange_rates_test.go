package handlers

import (
	"arf/currency-conversion/internal/database/connection"
	"arf/currency-conversion/internal/repository"
	"testing"

	"github.com/joho/godotenv"
)

func TestExchangeRateByCurrencyNames(t *testing.T) {

	if err := godotenv.Load("../../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	db := connection.Init()
	rep := repository.NewR(db)
	_, err := rep.GetExchangeRateByCurrencyNames("EUR", "TRY")

	if err != nil {
		t.Errorf("Exchange rate error = %v;", err.Error())
	}
}
