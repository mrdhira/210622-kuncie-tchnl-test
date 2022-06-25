package services

import (
	"context"

	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"
	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dto"
	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/repository"
)

type CartServices struct {
	repository repository.IRepository
}

// InitCartServices func
func InitCartServices(repository repository.IRepository) *CartServices {
	return &CartServices{
		repository: repository,
	}
}

// GetCarts func
func (s *CartServices) GetCarts(ctx context.Context, userID uint) (dto.Carts, error) {
	carts, err := s.repository.FindOrCreateCarts(ctx, dao.Carts{UserID: userID})
	if err != nil {
		return dto.Carts{}, err
	}

	carts.CartProducts, err = s.repository.FindAllCartProducts(ctx, carts.ID)
	if err != nil {
		return dto.Carts{}, err
	}

	var dtoCarts dto.Carts
	dtoCarts.DaoToDto(carts)

	return dtoCarts, nil
}

// CreateCartProducts func
func (s *CartServices) CreateCartProducts(ctx context.Context, userID uint, data dto.CreateCartProducts) (dto.Carts, error) {
	carts, err := s.repository.FindOrCreateCarts(ctx, dao.Carts{UserID: userID})
	if err != nil {
		return dto.Carts{}, err
	}

	_, err = s.repository.InsertOrUpdateCartProducts(
		ctx,
		dao.CartProducts{
			CartID:     carts.ID,
			ProductID:  data.ProductID,
			Quantity:   data.Quantity,
			Price:      data.Price,
			TotalPrice: float64(data.Quantity) * data.Price,
		},
	)
	if err != nil {
		return dto.Carts{}, err
	}

	carts.CartProducts, err = s.repository.FindAllCartProducts(ctx, carts.ID)
	if err != nil {
		return dto.Carts{}, err
	}

	var dtoCarts dto.Carts
	dtoCarts.DaoToDto(carts)

	return dtoCarts, nil
}

// UpdateCartProducts func
func (s *CartServices) UpdateCartProducts(ctx context.Context, userID uint, data dto.UpdateCartProducts) (dto.Carts, error) {
	carts, err := s.repository.FindOrCreateCarts(ctx, dao.Carts{UserID: userID})
	if err != nil {
		return dto.Carts{}, err
	}

	_, err = s.repository.InsertOrUpdateCartProducts(
		ctx,
		dao.CartProducts{
			CartID:     carts.ID,
			ProductID:  data.ProductID,
			Quantity:   data.Quantity,
			Price:      data.Price,
			TotalPrice: float64(data.Quantity) * data.Price,
		},
	)
	if err != nil {
		return dto.Carts{}, err
	}

	carts.CartProducts, err = s.repository.FindAllCartProducts(ctx, carts.ID)
	if err != nil {
		return dto.Carts{}, err
	}

	var dtoCarts dto.Carts
	dtoCarts.DaoToDto(carts)

	return dtoCarts, nil
}

// DeleteCartProducts func
func (s *CartServices) DeleteCartProducts(ctx context.Context, userID uint, data dto.DeleteCartProducts) error {
	carts, err := s.repository.FindOrCreateCarts(ctx, dao.Carts{UserID: userID})
	if err != nil {
		return err
	}

	err = s.repository.DeleteCartProducts(ctx, carts.ID, []uint{data.ProductID})
	if err != nil {
		return err
	}

	return nil
}
