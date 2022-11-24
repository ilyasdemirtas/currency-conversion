package repository

import (
	"arf/currency-conversion/api/internal/models"
	"fmt"

	"github.com/shopspring/decimal"
)

type UserWalletAccount struct {
	Balance decimal.Decimal
	Code    string
}

func (r R) WalletAccountByUser(userId uint32) ([]UserWalletAccount, error) {
	var data []UserWalletAccount

	result := r.db.Table("wallet_accounts as uwa").
		Joins("JOIN currencies as c ON c.id=uwa.currency_id").
		Select("uwa.balance, c.code").
		Where("uwa.user_id = ?", userId).
		Find(&data)

	if result.Error != nil {
		return nil, fmt.Errorf("user wallet account %v", result.Error)
	}

	return data, nil
}

func (r R) WalletAccounts() ([]models.WalletAccount, error) {
	var data []models.WalletAccount

	result := r.db.Model(&models.WalletAccount{}).Scan(&data)
	if result.Error != nil {
		return nil, fmt.Errorf("wallet accounts %v", result.Error)
	}

	return data, nil
}
