package config

import "arf/currency-conversion/api/config/helpers"

type APP struct {
	ApiSecret     string
	TokenLifeHour string
}

func NewApp() APP {
	return APP{
		ApiSecret:     helpers.Getenv("API_SECRET", "app_secret"),
		TokenLifeHour: helpers.Getenv("TOKEN_LIFE_HOUR", "1"),
	}
}
