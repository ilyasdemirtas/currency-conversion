package models

import "gorm.io/gorm"

type Currency struct {
	ID         uint32 `json:"id"`
	Name       string `gorm:"size:20;index:idx_name_code,unique" json:"name"`
	Code       string `gorm:"size:10;index:idx_name_code,unique" json:"code"`
	gorm.Model `json:"omitempty"`
}
