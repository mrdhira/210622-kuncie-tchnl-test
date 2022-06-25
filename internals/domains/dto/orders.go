package dto

import "github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"

type Orders struct {
	UserID        uint            `json:"user_id"`
	FinalAmount   float64         `json:"total_price"`
	OrderProducts []OrderProducts `json:"order_products"`
}

type OrderProducts struct {
	OrderID            uint          `json:"order_id"`
	ProductID          uint          `json:"product_id"`
	Quantity           int32         `json:"quantity"`
	Price              float64       `json:"price"`
	TotalPrice         float64       `json:"total_price"`
	FreeGiftQuantity   int32         `json:"free_gift_quantity"`
	DiscountPercentage float64       `json:"discount_percentage"`
	DiscountAmount     float64       `json:"discount_amount"`
	FinalAmount        float64       `json:"final_amount"`
	IsPromo            bool          `json:"is_promo"`
	PromoID            uint          `json:"promo_id"`
	PromoType          dao.PromoType `json:"promo_type"`
}

// daoToDto function converts a dao struct to dto struct
func (dto *Orders) DaoToDto(dao dao.Orders) {
	dto = &Orders{
		UserID:        dao.UserID,
		FinalAmount:   dao.FinalAmount,
		OrderProducts: make([]OrderProducts, 0),
	}
	for _, v := range dao.OrderProducts {
		dto.OrderProducts = append(dto.OrderProducts, OrderProducts{
			OrderID:            v.OrderID,
			ProductID:          v.ProductID,
			Quantity:           v.Quantity,
			Price:              v.Price,
			TotalPrice:         v.TotalPrice,
			FreeGiftQuantity:   v.FreeGiftQuantity,
			DiscountPercentage: v.DiscountPercentage,
			DiscountAmount:     v.DiscountAmount,
			FinalAmount:        v.FinalAmount,
			IsPromo:            v.IsPromo,
			PromoID:            v.PromoID,
			PromoType:          v.PromoType,
		})
	}
}
