package repository

import (
	"context"

	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"
	"gorm.io/gorm"
)

type ctxKey string

const trxKey = ctxKey("trx")

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

// getConnection func
func (r *Repository) getConnection(ctx context.Context) *gorm.DB {
	value := ctx.Value(trxKey).(*gorm.DB)
	if value != nil {
		return value.WithContext(ctx)
	}
	return r.client.WithContext(ctx)
}

// BeginTrx func
func (r *Repository) BeginTrx(ctx context.Context) context.Context {
	return context.WithValue(ctx, trxKey, r.client.Begin())
}

// CommitTrx func
func (r *Repository) CommitTrx(ctx context.Context) error {
	value := ctx.Value(trxKey).(*gorm.DB)
	return value.Commit().Error
}

// RollbackTrx func
func (r *Repository) RollbackTrx(ctx context.Context) error {
	value := ctx.Value(trxKey).(*gorm.DB)
	return value.Rollback().Error
}

// IRepository interface
type IRepository interface {
	BeginTrx(ctx context.Context) context.Context
	CommitTrx(ctx context.Context) error
	RollbackTrx(ctx context.Context) error

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
