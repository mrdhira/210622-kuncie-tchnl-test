package repository

import (
	"context"

	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"
	"gorm.io/gorm"
)

// Repository struct
type Repository struct {
	client *gorm.DB
}

// InitRepository func
func InitRepository(client *gorm.DB) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) Begin() *gorm.DB {
	return r.client.Begin()
}

func (r *Repository) Commit() *gorm.DB {
	return r.client.Commit()
}

func (r *Repository) Rollback() *gorm.DB {
	return r.client.Rollback()
}

// IRepository interface
type IRepository interface {
	Begin() *gorm.DB
	Commit() *gorm.DB
	Rollback() *gorm.DB

	// Carts
	FindOrCreateCarts(ctx context.Context, carts dao.Carts) (dao.Carts, error)
	FindAllCartProducts(ctx context.Context, cartID uint) ([]dao.CartProducts, error)
	InsertOrUpdateCartProducts(ctx context.Context, cartProducts dao.CartProducts) (dao.CartProducts, error)
	DeleteCartProducts(ctx context.Context, cartID uint, productIDs []uint) error

	// Promos
	GetPromos(ctx context.Context) ([]dao.Promos, error)

	// Orders
	CreateOrders(ctx context.Context, orders dao.Orders) (dao.Orders, error)

	// Products
	GetProducts(ctx context.Context, productID uint) (dao.Products, error)
	UpdateProductsQty(ctx context.Context, productID uint, qty int32) error
}
