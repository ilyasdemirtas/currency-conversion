package config

import "arf/currency-conversion/config/helpers"

type DB struct {
	Driver               string
	Host                 string
	Port                 string
	Database             string
	User                 string
	Password             string
	AllowNativePasswords string
	Timezone             string
}

func NewDB() DB {
	return DB{
		Driver:               helpers.Getenv("DB_DRIVER", "mysql"),
		Host:                 helpers.Getenv("DB_HOST", "127.0.0.1"),
		Port:                 helpers.Getenv("DB_PORT", "3306"),
		Database:             helpers.Getenv("DB_DATABASE", ""),
		User:                 helpers.Getenv("DB_USER", ""),
		Password:             helpers.Getenv("DB_PASS", ""),
		AllowNativePasswords: helpers.Getenv("DB_ALLOW_NATIVE_PASSWORDS", "True"),
		Timezone:             helpers.Getenv("DB_TIMEZONE", "Europe/Istanbul"),
	}
}
