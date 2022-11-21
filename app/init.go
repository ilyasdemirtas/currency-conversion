package app

import (
	"arf/currency-conversion/internal/database"
	"arf/currency-conversion/internal/database/connection"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

func DbConn() {
	db = connection.Init()
	database.SeedData(db)
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func GetDbConn() *gorm.DB {
	return db
}
