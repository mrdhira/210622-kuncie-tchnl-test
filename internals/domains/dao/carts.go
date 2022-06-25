package dao

import "gorm.io/gorm"

type Carts struct {
	gorm.Model
	UserID       uint           `gorm:"not null"`
	TotalPrice   float64        `gorm:"type:decimal(10,2);not null"`
	CartProducts []CartProducts `gorm:"foreignkey:CartID"`
}

type CartProducts struct {
	gorm.Model
	CartID     uint    `gorm:"not null"`
	ProductID  uint    `gorm:"not null"`
	Quantity   int32   `gorm:"type:int(10);not null"`
	Price      float64 `gorm:"type:decimal(10,2);not null"`
	TotalPrice float64 `gorm:"type:decimal(10,2);not null"`
}
