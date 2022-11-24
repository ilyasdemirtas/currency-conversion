package repository

import (
	"arf/currency-conversion/api/internal/models"
	"fmt"
)

func (r R) Currencies() ([]models.Currency, error) {
	var data []models.Currency

	result := r.db.Model(&models.Currency{}).Scan(&data)
	if result.Error != nil {
		return nil, fmt.Errorf("currency %v", result.Error)
	}

	return data, nil
}
