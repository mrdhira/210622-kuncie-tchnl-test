package repository

import (
	"context"
	"log"

	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"
)

// CreateOrders func
func (r *Repository) CreateOrders(ctx context.Context, order dao.Orders) (dao.Orders, error) {
	db := r.client.WithContext(ctx)
	err := db.Table("orders").Create(&order).Error
	if err != nil {
		log.Printf("error when create orders: %v\n", err)
		return dao.Orders{}, err
	}

	return order, nil
}
