package services

import (
	"context"

	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dto"
)

type ICartServices interface {
	GetCarts(ctx context.Context, userID uint) (dto.Carts, error)
	CreateCartProducts(ctx context.Context, userID uint, data dto.CreateCartProducts) (dto.Carts, error)
	UpdateCartProducts(ctx context.Context, userID uint, data dto.UpdateCartProducts) (dto.Carts, error)
	DeleteCartProducts(ctx context.Context, userID uint, data dto.DeleteCartProducts) error
}

type ICheckoutServices interface {
	Checkouts(ctx context.Context, data dto.Checkouts) (dto.Orders, error)
}
