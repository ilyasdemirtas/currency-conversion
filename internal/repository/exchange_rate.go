package repository

import (
	"arf/currency-conversion/internal/models"
	"fmt"

	"github.com/shopspring/decimal"
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

func (r ExchangeRateR) ExchangeRates() ([]models.ExchangeRate, error) {
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

type ExcangeRate struct {
	Price      decimal.Decimal
	MarkupRate decimal.Decimal
}

func (r ExchangeRateR) GetExchangeRateByCurrencyNames(baseCurrency, counterCurrency string) (ExcangeRate, error) {

	var exchangeRate ExcangeRate
	result := r.db.Table("exchange_rates as er").
		Joins("JOIN currencies as c1 ON c1.id=er.base_currency_id").
		Joins("JOIN currencies as c2 ON c2.id=er.counter_currency_id").
		Select("er.price, er.markup_rate").
		Where("c1.code = ?", baseCurrency).
		Where("c2.code = ?", counterCurrency).
		First(&exchangeRate)

	if result.Error != nil {
		return ExcangeRate{}, result.Error
	}

	return exchangeRate, nil
}
