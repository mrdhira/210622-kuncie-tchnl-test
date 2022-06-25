package dto

import "github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"

type CheckPromos struct {
	PromoID             uint          `json:"promo_id"`
	PromoType           dao.PromoType `json:"promo_type"`
	ProductID           uint          `json:"product_id"`
	MinimumProductQty   int32         `json:"minimum_product_qty"`
	CounterProductQty   int32         `json:"counter_product_qty"`
	FreeGiftProductID   uint          `json:"free_gift_product_id"`
	FreeGiftProductQty  int32         `json:"free_gift_product_qty"`
	FixedProductID      uint          `json:"fixed_product_id"`
	FixedGiftProductQty int32         `json:"fixed_gift_product_qty"`
	DiscountPercentage  float64       `json:"discount_percentage"`
	IsRedeemed          bool          `json:"is_redeemed"`
}
