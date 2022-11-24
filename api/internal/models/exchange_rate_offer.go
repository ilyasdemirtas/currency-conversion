package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ExchangeRateOffer struct {
	ID                uint32          `json:"id"`
	BaseCurrencyId    uint32          `json:"base_currency_id"`
	CounterCurrencyId uint32          `json:"counter_currency_id"`
	Price             decimal.Decimal `gorm:"type:decimal(10,4);" json:"price"`
	MarkupRate        decimal.Decimal `gorm:"type:decimal(10,4);" json:"markup_rate"`
	BaseCurrency      Currency        `gorm:"constraint:OnUpdate:SET NULL,OnDelete:CASCADE;" json:"omitempty"`
	CounterCurrency   Currency        `gorm:"constraint:OnUpdate:SET NULL,OnDelete:CASCADE;" json:"omitempty"`
	gorm.Model        `json:"omitempty"`
}
