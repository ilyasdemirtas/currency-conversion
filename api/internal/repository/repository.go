package repository

import (
	"gorm.io/gorm"
)

type R struct {
	db *gorm.DB
}

func NewR(db *gorm.DB) R {
	return R{db: db}
}
