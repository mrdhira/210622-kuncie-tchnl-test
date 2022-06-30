package repository

import (
	"context"
	"log"

	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"
)

// GetPromos func
func (r *Repository) GetPromos(ctx context.Context) ([]dao.Promos, error) {
	var promos []dao.Promos
	db := r.getConnection(ctx)
	err := db.Table("promos").Find(&promos).Error
	if err != nil {
		log.Printf("error when getting promos: %v\n", err)
		return []dao.Promos{}, err
	}

	return []dao.Promos{}, nil
}
