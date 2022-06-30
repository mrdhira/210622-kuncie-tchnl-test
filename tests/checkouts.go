package tests

import "github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dto"

var (
	CHECKOUTS_NO_PROMO = dto.Checkouts{
		UserID:           1,
		CartID:           1,
		CheckoutProducts: CHECKOUT_PRODUCTS_NO_PROMO,
	}
	CHECKOUT_PRODUCTS_NO_PROMO = []dto.CheckoutProducts{{
		CartProductID: 1,
		ProductID:     PRODUCT_RASPBERRY_PI.ID,
		Price:         PRODUCT_RASPBERRY_PI.Price,
		Quantity:      1,
	}}
	CHECKOUTS_MACBOOK_PRO_GET_RASPBERRY_PI = dto.Checkouts{
		UserID:           1,
		CartID:           1,
		CheckoutProducts: CHECKOUT_PRODUCTS_MACBOOK_PRO_GET_RASPBERRY_PI,
	}
	CHECKOUT_PRODUCTS_MACBOOK_PRO_GET_RASPBERRY_PI = []dto.CheckoutProducts{{
		CartProductID: 1,
		ProductID:     PRODUCT_MACBOOK.ID,
		Price:         PRODUCT_MACBOOK.Price,
		Quantity:      1,
	}}
	CHECKOUTS_GOOGLE_HOME_BY_2_GET_3 = dto.Checkouts{
		UserID:           1,
		CartID:           1,
		CheckoutProducts: CHECKOUT_PRODUCTS_GOOGLE_HOME_BY_2_GET_3,
	}
	CHECKOUT_PRODUCTS_GOOGLE_HOME_BY_2_GET_3 = []dto.CheckoutProducts{{
		CartProductID: 1,
		ProductID:     PRODUCT_GOOGLE_HOME.ID,
		Price:         PRODUCT_GOOGLE_HOME.Price,
		Quantity:      3,
	}}
	CHECKOUTS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT = dto.Checkouts{
		UserID:           1,
		CartID:           1,
		CheckoutProducts: CHECKOUT_PRODUCTS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT,
	}
	CHECKOUT_PRODUCTS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT = []dto.CheckoutProducts{{
		CartProductID: 1,
		ProductID:     PRODUCT_ALEXA_SPEAKER.ID,
		Price:         PRODUCT_ALEXA_SPEAKER.Price,
		Quantity:      3,
	}}
)
