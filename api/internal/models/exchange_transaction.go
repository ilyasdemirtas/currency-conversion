package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ExchangeTransaction struct {
	ID                uint32          `json:"id"`
	UserId            uint32          `json:"user_id"`
	BaseCurrencyId    uint32          `json:"base_currency_id"`
	CounterCurrencyId uint32          `json:"counter_currency_id"`
	Price             decimal.Decimal `gorm:"type:decimal(10,4);" json:"price"`
	MarkupRate        decimal.Decimal `gorm:"type:decimal(10,4);" json:"markup_rate"`
	User              User            `gorm:"constraint:OnUpdate:SET NULL,OnDelete:CASCADE;"`
	BaseCurrency      Currency        `gorm:"constraint:OnUpdate:SET NULL,OnDelete:CASCADE;"`
	CounterCurrency   Currency        `gorm:"constraint:OnUpdate:SET NULL,OnDelete:CASCADE;"`
	gorm.Model
}
