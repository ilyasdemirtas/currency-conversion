package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type WalletAccount struct {
	ID         uint32          `json:"id"`
	UserId     uint32          `json:"user_id"`
	CurrencyId uint32          `json:"currency_id"`
	Balance    decimal.Decimal `gorm:"type:decimal(10,4);" json:"balance"`
	User       User            `gorm:"constraint:OnUpdate:SET NULL,OnDelete:CASCADE;" json:"omitempty"`
	Currency   Currency        `gorm:"constraint:OnUpdate:SET NULL,OnDelete:CASCADE;" json:"omitempty"`
	gorm.Model `json:"omitempty"`
}
