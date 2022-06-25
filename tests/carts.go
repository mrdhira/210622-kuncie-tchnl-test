package tests

import (
	"time"

	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"
	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dto"
	"gorm.io/gorm"
)

var (
	BASE_GORM_MODEL = gorm.Model{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	CARTS = dao.Carts{
		Model:        BASE_GORM_MODEL,
		UserID:       1,
		TotalPrice:   100,
		CartProducts: []dao.CartProducts{},
	}
	CART_PRODUCTS = []dao.CartProducts{{
		Model:      BASE_GORM_MODEL,
		CartID:     BASE_GORM_MODEL.ID,
		ProductID:  1,
		Quantity:   1,
		Price:      100,
		TotalPrice: 100,
	}}
	CART_PRODUCT = dao.CartProducts{
		Model:      BASE_GORM_MODEL,
		CartID:     BASE_GORM_MODEL.ID,
		ProductID:  1,
		Quantity:   1,
		Price:      100,
		TotalPrice: 100,
	}
	CREATE_CART_PRODUCTS = dto.CreateCartProducts{
		ProductID: 1,
		Quantity:  1,
		Price:     100,
	}
	UPDATE_CART_PRODUCTS = dto.UpdateCartProducts{
		ProductID: 1,
		Quantity:  1,
		Price:     100,
	}
	DELETE_CART_PRODUCT = dto.DeleteCartProducts{
		ProductID: 1,
	}
)
