package dao

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	SKU      string  `gorm:"type:varchar(100);not null"`
	Name     string  `gorm:"type:varchar(100);not null"`
	Price    float64 `gorm:"type:decimal(10,2);not null"`
	Quantity int32   `gorm:"type:int(10);not null"`
}
