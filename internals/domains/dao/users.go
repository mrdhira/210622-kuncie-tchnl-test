package dao

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);not null"`
}
