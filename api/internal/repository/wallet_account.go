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

func (r R) CheckUserBalanceByCurrency(userId, currencyId uint32, totalPrice decimal.Decimal) int {
	result := r.db.Model(models.WalletAccount{}).
		Where("user_id = ? AND currency_id = ? AND balance > ?", userId, currencyId, totalPrice).
		First(&models.WalletAccount{})
	return int(result.RowsAffected)
}

func (r R) WalletAccountByUserAndCurrency(userId, currencyId uint32) (models.WalletAccount, error) {
	var data models.WalletAccount

	err := r.db.Model(models.WalletAccount{}).
		Where("user_id = ? AND currency_id = ?", userId, currencyId).
		First(&data).
		Error

	if err != nil {
		return data, fmt.Errorf("user wallet account %v", err)
	}

	return data, nil
}

func (r R) UpdateUserWallet(data models.WalletAccount) error {
	err := r.db.Save(&data).Error
	if err != nil {
		return fmt.Errorf("user wallet account %v", err.Error())
	}
	return nil
}

func (r R) WalletAccounts() ([]models.WalletAccount, error) {
	var data []models.WalletAccount

	result := r.db.Model(&models.WalletAccount{}).Scan(&data)
	if result.Error != nil {
		return nil, fmt.Errorf("wallet accounts %v", result.Error)
	}

	return data, nil
}
