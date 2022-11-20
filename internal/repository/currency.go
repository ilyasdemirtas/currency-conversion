package repository

import (
	"arf/currency-conversion/internal/models"
	"fmt"
	"gorm.io/gorm"
)

type CurrencyR struct {
	db *gorm.DB
}

func NewCurrency(db *gorm.DB) CurrencyR {
	return CurrencyR{
		db: db,
	}
}

func (r CurrencyR) All() ([]models.Currency, error) {
	var currencies []models.Currency

	data := r.db.Model(&models.Currency{}).Scan(&currencies)
	if data.Error != nil {
		return nil, fmt.Errorf("Currency %v", data.Error)
	}

	return currencies, nil
}
