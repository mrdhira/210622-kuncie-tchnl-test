package services

import (
	"context"
	"errors"

	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"
	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dto"
	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/repository"
)

type CheckoutServices struct {
	repository repository.IRepository
}

// InitCheckoutServices func
func InitCheckoutServices(repository repository.IRepository) *CheckoutServices {
	return &CheckoutServices{
		repository: repository,
	}
}

// checkPromos func
func (s *CheckoutServices) checkPromos(ctx context.Context, orders dao.Orders) (dao.Orders, error) {
	// Get All Active Promos
	promos, err := s.repository.GetPromos(ctx)
	if err != nil {
		return orders, err
	}

	promoProductID := make(map[uint][]dto.CheckPromos, 0)
	for _, promo := range promos {
		switch promo.Type {
		case dao.PromoTypeFreeGift:
			promoProductID[promo.Threshold.ProductID] = append(promoProductID[promo.Threshold.ProductID], dto.CheckPromos{
				PromoID:            promo.ID,
				PromoType:          promo.Type,
				ProductID:          promo.Threshold.ProductID,
				MinimumProductQty:  promo.Threshold.MinimumProductQty,
				CounterProductQty:  0,
				FreeGiftProductID:  promo.Gift.FreeGiftProductID,
				FreeGiftProductQty: promo.Gift.FreeGiftProductQty,
				IsRedeemed:         false,
			})

		case dao.PromoTypeFixed:
			promoProductID[promo.Threshold.ProductID] = append(promoProductID[promo.Threshold.ProductID], dto.CheckPromos{
				PromoID:            promo.ID,
				PromoType:          promo.Type,
				ProductID:          promo.Threshold.ProductID,
				MinimumProductQty:  promo.Threshold.MinimumProductQty,
				CounterProductQty:  0,
				FreeGiftProductID:  promo.Gift.FixedProductID,
				FreeGiftProductQty: promo.Gift.FixedGiftProductQty,
				IsRedeemed:         false,
			})

		case dao.PromoTypeDiscount:
			promoProductID[promo.Threshold.ProductID] = append(promoProductID[promo.Threshold.ProductID], dto.CheckPromos{
				PromoID:            promo.ID,
				PromoType:          promo.Type,
				ProductID:          promo.Threshold.ProductID,
				MinimumProductQty:  promo.Threshold.MinimumProductQty,
				CounterProductQty:  0,
				DiscountPercentage: promo.Gift.DiscountPercentage,
				IsRedeemed:         false,
			})

		default:
			continue
		}
	}

	for idxOrderProducts, orderProducts := range orders.OrderProducts {
		// Check if order product id is in promoProductID map
		if _, ok := promoProductID[orderProducts.ProductID]; ok {
			for idx := range promoProductID[orderProducts.ProductID] {
				promoProductID[orderProducts.ProductID][idx].CounterProductQty = orderProducts.Quantity

				// check if prder product quantity more than threshold,
				// and promos not yet redeemed
				// and orders product not yet applied by any promos
				// then apply promos
				if promoProductID[orderProducts.ProductID][idx].CounterProductQty >= promoProductID[orderProducts.ProductID][idx].MinimumProductQty &&
					!promoProductID[orderProducts.ProductID][idx].IsRedeemed &&
					!orders.OrderProducts[idxOrderProducts].IsPromo {
					switch promoProductID[orderProducts.ProductID][idx].PromoType {
					case dao.PromoTypeFreeGift:
						orders.OrderProducts = append(orders.OrderProducts, dao.OrderProducts{
							OrderID:            orders.ID,
							ProductID:          promoProductID[orderProducts.ProductID][idx].FreeGiftProductID,
							Quantity:           promoProductID[orderProducts.ProductID][idx].FreeGiftProductQty,
							Price:              0,
							TotalPrice:         0,
							FreeGiftQuantity:   0,
							DiscountPercentage: 0,
							DiscountAmount:     0,
							FinalAmount:        0,
							IsPromo:            true,
							PromoID:            promoProductID[orderProducts.ProductID][idx].PromoID,
							PromoType:          promoProductID[orderProducts.ProductID][idx].PromoType,
						})

						promoProductID[orderProducts.ProductID][idx].IsRedeemed = true
					case dao.PromoTypeFixed:
						orders.OrderProducts[idxOrderProducts].FreeGiftQuantity = promoProductID[orderProducts.ProductID][idx].FreeGiftProductQty
						orders.OrderProducts[idxOrderProducts].FinalAmount = orders.OrderProducts[idxOrderProducts].Price * (float64(orders.OrderProducts[idxOrderProducts].Quantity) - float64(promoProductID[orderProducts.ProductID][idx].FreeGiftProductQty))
						orders.OrderProducts[idxOrderProducts].IsPromo = true
						orders.OrderProducts[idxOrderProducts].PromoID = promoProductID[orderProducts.ProductID][idx].PromoID
						orders.OrderProducts[idxOrderProducts].PromoType = promoProductID[orderProducts.ProductID][idx].PromoType

						promoProductID[orderProducts.ProductID][idx].IsRedeemed = true
					case dao.PromoTypeDiscount:
						orders.OrderProducts[idxOrderProducts].DiscountPercentage = promoProductID[orderProducts.ProductID][idx].DiscountPercentage
						orders.OrderProducts[idxOrderProducts].DiscountAmount = orders.OrderProducts[idxOrderProducts].TotalPrice * float64(promoProductID[orderProducts.ProductID][idx].DiscountPercentage) / 100
						orders.OrderProducts[idxOrderProducts].FinalAmount = orders.OrderProducts[idxOrderProducts].TotalPrice - orders.OrderProducts[idxOrderProducts].DiscountAmount
						orders.OrderProducts[idxOrderProducts].IsPromo = true
						orders.OrderProducts[idxOrderProducts].PromoID = promoProductID[orderProducts.ProductID][idx].PromoID
						orders.OrderProducts[idxOrderProducts].PromoType = promoProductID[orderProducts.ProductID][idx].PromoType

						promoProductID[orderProducts.ProductID][idx].IsRedeemed = true
					default:
						continue
					}
				}
			}
		}

		// Calculate final amount of orders from order products
		orders.FinalAmount += orders.OrderProducts[idxOrderProducts].FinalAmount
	}

	return orders, nil
}

// createOrder func
func (s *CheckoutServices) createOrder(ctx context.Context, orders dao.Orders) (dao.Orders, error) {
	repositoryTrans := repository.InitRepository(s.repository.Begin())

	var err error
	defer func(err error) {
		if err != nil {
			repositoryTrans.Rollback()
		} else {
			repositoryTrans.Commit()
		}
	}(err)

	// check if ordered product quantity is valid with master product quantity
	for _, orderProducts := range orders.OrderProducts {
		product, err := repositoryTrans.GetProducts(ctx, orderProducts.ProductID)
		if err != nil {
			return dao.Orders{}, err
		}

		// if not matched, return error
		if product.Quantity != orderProducts.Quantity {
			return dao.Orders{}, errors.New("product quantity not match")
		} else {
			// if matched, update master product quantity
			err := repositoryTrans.UpdateProductsQty(ctx, orderProducts.ProductID, product.Quantity-orderProducts.Quantity)
			if err != nil {
				return dao.Orders{}, err
			}
		}
	}

	orders, err = s.repository.CreateOrders(ctx, orders)
	if err != nil {
		return dao.Orders{}, err
	}

	return orders, nil
}

// Checkouts func
func (s *CheckoutServices) Checkouts(ctx context.Context, data dto.Checkouts) (dto.Orders, error) {
	orders := dao.Orders{
		UserID:        data.UserID,
		FinalAmount:   0,
		OrderProducts: []dao.OrderProducts{},
	}

	var productIDs []uint
	for _, orderProducts := range data.CheckoutProducts {
		orders.OrderProducts = append(orders.OrderProducts, dao.OrderProducts{
			ProductID:          orderProducts.ProductID,
			Quantity:           orderProducts.Quantity,
			Price:              orderProducts.Price,
			TotalPrice:         float64(orderProducts.Quantity) * orderProducts.Price,
			FreeGiftQuantity:   0,
			DiscountPercentage: 0,
			DiscountAmount:     0,
			FinalAmount:        float64(orderProducts.Quantity) * orderProducts.Price,
			IsPromo:            false,
			PromoID:            0,
			PromoType:          "",
		})

		productIDs = append(productIDs, orderProducts.ProductID)
	}

	var err error
	orders, err = s.checkPromos(ctx, orders)
	if err != nil {
		return dto.Orders{}, err
	}

	orders, err = s.createOrder(ctx, orders)
	if err != nil {
		return dto.Orders{}, err
	}

	// clear cart products that has been checkout
	err = s.repository.DeleteCartProducts(ctx, data.CartID, productIDs)
	if err != nil {
		return dto.Orders{}, err
	}

	var dtoOrders dto.Orders
	dtoOrders.DaoToDto(orders)

	return dtoOrders, nil
}
