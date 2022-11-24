package repository

import (
	"arf/currency-conversion/api/internal/models"
	"fmt"
	"time"
)

func (r R) CreateExchangeOffer(data models.ExchangeRateOffer) (models.ExchangeRateOffer, error) {
	result := r.db.Create(&data)

	if result.Error != nil {
		return data, fmt.Errorf("exchange rate offer %v", result.Error)
	}

	return data, nil
}

func (r R) GetExchangeRateOfferByCurrencyNames(baseCurrency, counterCurrency string) (
	models.ExchangeRateOffer, error) {

	var data models.ExchangeRateOffer
	result := r.db.Table("exchange_rate_offers as ero").
		Select("ero.base_currency_id, ero.counter_currency_id, ero.price, ero.markup_rate").
		Joins("INNER JOIN currencies as c1 ON c1.id = ero.base_currency_id").
		Joins("INNER JOIN currencies as c2 ON c2.id = ero.counter_currency_id").
		Where("c1.code = ?", baseCurrency).
		Where("c2.code = ?", counterCurrency).
		Where("ero.created_at >=? ", time.Now().Add(time.Duration(-3)*time.Minute)).
		Last(&data)

	if result.Error != nil {
		return data, fmt.Errorf("exchage rate offer %v", result.Error)
	}

	return data, nil
}
