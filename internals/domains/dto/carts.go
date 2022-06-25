package dto

import "github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"

type Carts struct {
	UserID       uint           `json:"user_id"`
	TotalPrice   float64        `json:"total_price"`
	CartProducts []CartProducts `json:"cart_products"`
}

type CartProducts struct {
	ProductID  uint    `json:"product_id" binding:"required"`
	Quantity   int32   `json:"quantity" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	TotalPrice float64 `json:"total_price" binding:"required"`
}

type CreateCartProducts struct {
	ProductID uint    `json:"product_id"`
	Quantity  int32   `json:"quantity"`
	Price     float64 `json:"price"`
}

type UpdateCartProducts struct {
	ProductID uint    `json:"product_id"`
	Quantity  int32   `json:"quantity"`
	Price     float64 `json:"price"`
}

type DeleteCartProducts struct {
	ProductID uint `json:"product_id"`
}

// CartsDaoToDto function converts a dao struct to dto struct
func CartsDaoToDto(dao dao.Carts) Carts {
	dto := Carts{
		UserID:       dao.UserID,
		TotalPrice:   dao.TotalPrice,
		CartProducts: make([]CartProducts, 0),
	}
	for _, v := range dao.CartProducts {
		dto.CartProducts = append(dto.CartProducts, CartProducts{
			ProductID:  v.ProductID,
			Quantity:   v.Quantity,
			Price:      v.Price,
			TotalPrice: v.TotalPrice,
		})
	}

	return dto
}
