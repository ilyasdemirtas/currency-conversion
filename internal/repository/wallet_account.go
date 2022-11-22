package repository

import (
	"arf/currency-conversion/internal/models"
	"fmt"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type UserWalletAccount struct {
	Balance decimal.Decimal
	Code    string
}

type WalletAccountR struct {
	db *gorm.DB
}

func NewWalletAccount(db *gorm.DB) WalletAccountR {
	return WalletAccountR{
		db: db,
	}
}

func (r WalletAccountR) WalletAccountByUser(userId uint32) ([]UserWalletAccount, error) {
	var walletAccounts []UserWalletAccount

	data := r.db.Table("wallet_accounts as uwa").
		Joins("JOIN currencies as c ON c.id=uwa.currency_id").
		Select("uwa.balance, c.code").
		Where("uwa.user_id = ?", userId).
		Find(&walletAccounts)

	if data.Error != nil {
		return nil, fmt.Errorf("user wallet account %v", data.Error)
	}

	return walletAccounts, nil
}

func (r WalletAccountR) WalletAccounts() ([]models.WalletAccount, error) {
	var walletAccounts []models.WalletAccount

	data := r.db.Model(&models.WalletAccount{}).Scan(&walletAccounts)
	if data.Error != nil {
		return nil, fmt.Errorf("wallet accounts %v", data.Error)
	}

	return walletAccounts, nil
}
