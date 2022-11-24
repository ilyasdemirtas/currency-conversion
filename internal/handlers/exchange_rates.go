package handlers

import (
	"arf/currency-conversion/app"
	"arf/currency-conversion/internal/models"
	"arf/currency-conversion/internal/repository"
)

func GetExchangeRateByCurrencyNames(baseCurrency, counterCurrency string) (models.ExchangeRate, error) {

	rep := repository.NewR(app.GetDbConn())
	data, err := rep.GetExchangeRateByCurrencyNames(baseCurrency, counterCurrency)

	if err != nil {
		return models.ExchangeRate{}, err
	}

	return data, nil
}
