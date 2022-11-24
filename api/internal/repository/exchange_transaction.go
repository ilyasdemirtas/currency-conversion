package repository

import (
	"arf/currency-conversion/api/internal/models"
	"fmt"
)

func (r R) CreateExchangeTransaction(data models.ExchangeTransaction) (bool, error) {
	err := r.db.Create(&data).Error

	if err != nil {
		return false, fmt.Errorf("exchange transaction %v", err)
	}

	return true, nil
}
