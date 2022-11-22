package database

import (
	"arf/currency-conversion/internal/models"
	"errors"

	"github.com/shopspring/decimal"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) {

	if db.Migrator().HasTable(&models.User{}) {
		addUser(db)
	}

	if db.Migrator().HasTable(&models.Currency{}) {
		addCurrencies(db)
	}

	if db.Migrator().HasTable(&models.ExchangeRate{}) {
		addExchangeRates(db)
	}

	if db.Migrator().HasTable(&models.WalletAccount{}) {
		addWalletAccount(db)
	}
}

func addUser(db *gorm.DB) {

	if err := db.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123qwe"), bcrypt.DefaultCost)

		pass := string(hashedPassword)
		user := models.User{Username: "ilyasdemirtas", Email: "ilyasdemirtas@hotmail.com.tr", Password: pass}

		db.Create(&user)

	}

}

func addCurrencies(db *gorm.DB) {

	if err := db.First(&models.Currency{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {

		currencies := []models.Currency{
			{Name: "Turkish Lira", Code: "TRY"},
			{Name: "Dollar", Code: "USD"},
			{Name: "Euro", Code: "EUR"},
		}

		db.Create(&currencies)

	}

}

func addWalletAccount(db *gorm.DB) {

	if err := db.First(&models.WalletAccount{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {

		balances := []models.WalletAccount{
			{UserId: 1, CurrencyId: 1, Balance: decimal.NewFromFloat(1000)},
			{UserId: 1, CurrencyId: 2, Balance: decimal.NewFromFloat(1000)},
			{UserId: 1, CurrencyId: 3, Balance: decimal.NewFromFloat(1000)},
		}

		db.Create(&balances)
	}

}

func addExchangeRates(db *gorm.DB) {

	if err := db.First(&models.ExchangeRate{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {

		exchangeRates := []models.ExchangeRate{
			{BaseCurrencyId: 2, CounterCurrencyId: 1, Price: decimal.NewFromFloat(18.5383), MarkupRate: decimal.NewFromFloat(0.40)},
			{BaseCurrencyId: 3, CounterCurrencyId: 1, Price: decimal.NewFromFloat(19.3890), MarkupRate: decimal.NewFromFloat(0.50)},
			{BaseCurrencyId: 3, CounterCurrencyId: 2, Price: decimal.NewFromFloat(1.1234), MarkupRate: decimal.NewFromFloat(0.60)},
		}

		db.Create(&exchangeRates)

	}

}
