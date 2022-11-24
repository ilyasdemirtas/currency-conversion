package repository

import (
	"arf/currency-conversion/api/internal/models"
	"fmt"
)

func (r R) ExchangeRates() ([]models.ExchangeRate, error) {
	var data []models.ExchangeRate

	result := r.db.Model(&models.ExchangeRate{}).Scan(&data)
	if result.Error != nil {
		return nil, fmt.Errorf("ExcangeRate %v", result.Error)
	}

	return data, nil
}

func (r R) ExchangeRateByBaseCurrencyIdAndCounterCurrencyId(baseCurrencyId int, counterCurrencyId int) (
	[]models.ExchangeRate,
	error,
) {
	var data []models.ExchangeRate

	result := r.db.Model(&models.ExchangeRate{}).
		Where("base_currency_id = ? AND counter_currency_id = ?", baseCurrencyId, counterCurrencyId).
		Scan(&data)

	if result.Error != nil {
		return nil, fmt.Errorf("ExcangeRate %v", result.Error)
	}

	return data, nil
}

func (r R) GetExchangeRateByCurrencyNames(baseCurrency, counterCurrency string) (
	models.ExchangeRate, error) {

	var data models.ExchangeRate

	result := r.db.Table("exchange_rates as er").
		Joins("JOIN currencies as c1 ON c1.id=er.base_currency_id").
		Joins("JOIN currencies as c2 ON c2.id=er.counter_currency_id").
		Select("er.base_currency_id, er.counter_currency_id, er.price, er.markup_rate").
		Where("c1.code = ?", baseCurrency).
		Where("c2.code = ?", counterCurrency).
		First(&data)

	if result.Error != nil {
		return data, fmt.Errorf("exchange rate offer %v", result.Error)
	}

	return data, nil
}
