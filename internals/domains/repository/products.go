package repository

import (
	"context"
	"log"

	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"
)

// GetProducts func
func (r *Repository) GetProducts(ctx context.Context, productID uint) (dao.Products, error) {
	var product dao.Products
	db := r.getConnection(ctx)
	err := db.Table("products").First(&product, productID).Error
	if err != nil {
		log.Printf("error when getting product: %v\n", err)
		return dao.Products{}, err
	}

	return dao.Products{}, nil
}

// UpdateProductsQty func
func (r *Repository) UpdateProductsQty(ctx context.Context, productID uint, qty int32) error {
	db := r.getConnection(ctx)
	err := db.Table("products").Where("id = ?", productID).Update("qty", qty).Error
	if err != nil {
		log.Printf("error when updating product qty: %v\n", err)
		return err
	}

	return nil
}
