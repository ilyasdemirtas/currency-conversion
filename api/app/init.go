package app

import (
	"arf/currency-conversion/api/internal/database"
	"arf/currency-conversion/api/internal/database/connection"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	loadEnv()
	dbConn()
}

func dbConn() {
	db = connection.Init()
	database.SeedData(db)
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func GetDbConn() *gorm.DB {
	return db
}
