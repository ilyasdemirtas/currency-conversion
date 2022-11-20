package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type UserWalletAccount struct {
	ID         uint32          `json:"id"`
	UserId     uint32          `json:"user_id"`
	CurrencyId uint32          `json:"currency_id"`
	Balance    decimal.Decimal `gorm:"type:decimal(10,4);" json:"balance"`
	User       User            `gorm:"constraint:OnUpdate:SET NULL,OnDelete:CASCADE;"`
	Currency   Currency        `gorm:"constraint:OnUpdate:SET NULL,OnDelete:CASCADE;"`
	gorm.Model
}
