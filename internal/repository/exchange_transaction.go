package repository

import (
	"arf/currency-conversion/internal/models"
	"fmt"
)

func (r R) CreateExchangeTransaction(data models.ExchangeTransaction) (models.ExchangeTransaction, error) {
	result := r.db.Create(&data)

	if result.Error != nil {
		return data, fmt.Errorf("exchange transaction %v", result.Error)
	}

	return data, nil
}
