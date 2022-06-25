package dao

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	UserID        uint            `gorm:"not null"`
	FinalAmount   float64         `gorm:"type:decimal(10,2);not null"`
	OrderProducts []OrderProducts `gorm:"foreignkey:OrderID"`
}

type OrderProducts struct {
	gorm.Model
	OrderID            uint      `gorm:"not null"`
	ProductID          uint      `gorm:"not null"`
	Quantity           int32     `gorm:"type:int(10);not null"`
	Price              float64   `gorm:"type:decimal(10,2);not null"`
	TotalPrice         float64   `gorm:"type:decimal(10,2);not null"`
	FreeGiftQuantity   int32     `gorm:"type:int(10);not null"`
	DiscountPercentage float64   `gorm:"type:decimal(10,2);not null"`
	DiscountAmount     float64   `gorm:"type:decimal(10,2);not null"`
	FinalAmount        float64   `gorm:"type:decimal(10,2);not null"`
	IsPromo            bool      `gorm:"not null"`
	PromoID            uint      `gorm:"not null"`
	PromoType          PromoType `gorm:"type:varchar(100);not null"`
}
