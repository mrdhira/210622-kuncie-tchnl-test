package dao

import "gorm.io/gorm"

type PromoType string

const (
	PromoTypeFreeGift PromoType = "free_gift" // Get something free when threshold meet
	PromoTypeFixed    PromoType = "fixed"     // Get a free fixed things/money when threshold meet
	PromoTypeDiscount PromoType = "discount"  // Get a percentage of discount when threshold meet
)

type PromoThreshold struct {
	ProductID         uint  `gorm:"not null"`
	MinimumProductQty int32 `gorm:"type:int(10);not null"`
}

type PromoGift struct {
	FreeGiftProductID   uint    `gorm:"not null"`                    // For promo type free_gift
	FreeGiftProductQty  int32   `gorm:"type:int(10);not null"`       // For promo type free_gift
	FixedProductID      uint    `gorm:"not null"`                    // For promo type fixed
	FixedGiftProductQty int32   `gorm:"type:int(10);not null"`       // For promo type fixed
	DiscountPercentage  float64 `gorm:"type:decimal(10,2);not null"` // For promo type discount
}

type Promos struct {
	gorm.Model
	Type      PromoType      `gorm:"type:varchar(100);not null"`
	Threshold PromoThreshold `gorm:"type:jsonb"`
	Gift      PromoGift      `gorm:"type:jsonb"`
}
