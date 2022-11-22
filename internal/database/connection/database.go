package connection

import (
	"arf/currency-conversion/config/helpers"
	"arf/currency-conversion/internal/database"
	"arf/currency-conversion/internal/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func Init() (db *gorm.DB) {
	var err error
	db, err = gorm.Open(database.DbDriver(), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database ", helpers.Getenv("DB_DRIVER", ""))
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", helpers.Getenv("DB_DRIVER", ""))
	}

	db.AutoMigrate(
		&models.User{},
		&models.Currency{},
		&models.ExchangeRate{},
		&models.ExchangeRateOffer{},
		&models.WalletAccount{},
	)

	return db
}
