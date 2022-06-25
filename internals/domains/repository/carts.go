package repository

import (
	"context"
	"log"

	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"
	"gorm.io/gorm/clause"
)

// FindOrCreateCarts func
func (r *Repository) FindOrCreateCarts(ctx context.Context, carts dao.Carts) (dao.Carts, error) {
	db := r.client.WithContext(ctx)
	err := db.Table("carts").FirstOrCreate(&carts, dao.Carts{UserID: carts.UserID}).Error
	if err != nil {
		log.Printf("error when first or create cart: %v\n", err)
		return dao.Carts{}, err
	}

	return carts, nil
}

// FindAllCartProducts func
func (r *Repository) FindAllCartProducts(ctx context.Context, cartID uint) ([]dao.CartProducts, error) {
	var cartProducts []dao.CartProducts
	db := r.client.WithContext(ctx)
	err := db.Table("cart_products").Where("cart_id = ?", cartID).Find(&cartProducts).Error
	if err != nil {
		log.Printf("error when getting cart products: %v\n", err)
		return []dao.CartProducts{}, err
	}

	return cartProducts, nil
}

// InsertOrUpdateCartProducts func
func (r *Repository) InsertOrUpdateCartProducts(ctx context.Context, cartProducts dao.CartProducts) (dao.CartProducts, error) {
	db := r.client.WithContext(ctx)
	err := db.Table("cart_products").Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "cart_id"}, {Name: "product_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"quantity", "price", "total_price"}),
		},
	).Create(&cartProducts).Error
	if err != nil {
		log.Printf("error when insert or update cart products: %v\n", err)
		return dao.CartProducts{}, err
	}

	return cartProducts, nil
}

// DeleteCartProducts func
func (r *Repository) DeleteCartProducts(ctx context.Context, cart_id uint, productIDs []uint) error {
	db := r.client.WithContext(ctx)
	err := db.
		Table("cart_products").
		Where("cart_id = ?", cart_id).
		Where("product_id IN ?", productIDs).
		Delete(&dao.CartProducts{}).
		Error
	if err != nil {
		log.Printf("error when delete cart products: %v\n", err)
		return err
	}

	return nil
}
