package repository

import (
	"arf/currency-conversion/internal/models"
	"fmt"
	"gorm.io/gorm"
)

type ExchangeRateR struct {
	db *gorm.DB
}

func NewExchangeRate(db *gorm.DB) ExchangeRateR {
	return ExchangeRateR{
		db: db,
	}
}

func (r ExchangeRateR) All() ([]models.ExchangeRate, error) {
	var exchangeRates []models.ExchangeRate

	data := r.db.Model(&models.ExchangeRate{}).Scan(&exchangeRates)
	if data.Error != nil {
		return nil, fmt.Errorf("ExcangeRate %v", data.Error)
	}

	if data.RowsAffected == 0 {
		return nil, fmt.Errorf("not found Exchange Rates")
	}

	return exchangeRates, nil
}

func (r ExchangeRateR) ExchangeRateByBaseCurrencyIdAndCounterCurrencyId(baseCurrencyId int, counterCurrencyId int) (
	[]models.ExchangeRate,
	error,
) {
	var exchangeRates []models.ExchangeRate

	data := r.db.Model(&models.ExchangeRate{}).
		Where("base_currency_id = ? AND counter_currency_id = ?", baseCurrencyId, counterCurrencyId).
		Scan(&exchangeRates)

	if data.Error != nil {
		return nil, fmt.Errorf("ExcangeRate %v", data.Error)
	}

	return exchangeRates, nil
}
