package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID            uint32          `json:"id"`
	Username      string          `gorm:"size:50;index:idx_name_email,unique" json:"username"`
	Email         string          `gorm:"size:255;index:idx_name_email,unique" json:"email"`
	Password      string          `gorm:"size:255" json:"password"`
	WalletAccount []WalletAccount `json:"omitempty"`
	gorm.Model    `json:"omitempty"`
}
