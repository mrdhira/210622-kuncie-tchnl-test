package tests

import "github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"

var (
	ORDERS_NO_PROMO = dao.Orders{
		Model:         BASE_GORM_MODEL,
		UserID:        1,
		FinalAmount:   PRODUCT_RASPBERRY_PI.Price,
		OrderProducts: []dao.OrderProducts{},
	}
	ORDER_PRODUCTS_NO_PROMO = []dao.OrderProducts{{
		Model:              BASE_GORM_MODEL,
		OrderID:            BASE_GORM_MODEL.ID,
		ProductID:          PRODUCT_RASPBERRY_PI.ID,
		Quantity:           1,
		Price:              PRODUCT_RASPBERRY_PI.Price,
		TotalPrice:         PRODUCT_RASPBERRY_PI.Price,
		FreeGiftQuantity:   0,
		DiscountPercentage: 0,
		DiscountAmount:     0,
		FinalAmount:        PRODUCT_RASPBERRY_PI.Price,
		IsPromo:            false,
		PromoID:            0,
		PromoType:          "",
	}}
	ORDERS_MACBOOK_PRO_GET_RASPBERRY_PI = dao.Orders{
		Model:         BASE_GORM_MODEL,
		UserID:        1,
		FinalAmount:   PRODUCT_MACBOOK.Price,
		OrderProducts: ORDER_PRODUCTS_MACBOOK_PRO_GET_RASPBERRY_PI,
	}
	ORDER_PRODUCTS_MACBOOK_PRO_GET_RASPBERRY_PI = []dao.OrderProducts{
		{
			Model:              BASE_GORM_MODEL,
			OrderID:            BASE_GORM_MODEL.ID,
			ProductID:          PRODUCT_MACBOOK.ID,
			Quantity:           1,
			Price:              PRODUCT_MACBOOK.Price,
			TotalPrice:         PRODUCT_MACBOOK.Price,
			FreeGiftQuantity:   0,
			DiscountPercentage: 0,
			DiscountAmount:     0,
			FinalAmount:        PRODUCT_MACBOOK.Price,
			IsPromo:            false,
			PromoID:            0,
			PromoType:          "",
		},
		{
			Model:              BASE_GORM_MODEL,
			OrderID:            BASE_GORM_MODEL.ID,
			ProductID:          PRODUCT_RASPBERRY_PI.ID,
			Quantity:           1,
			Price:              0,
			TotalPrice:         0,
			FreeGiftQuantity:   0,
			DiscountPercentage: 0,
			DiscountAmount:     0,
			FinalAmount:        0,
			IsPromo:            true,
			PromoID:            PROMO_MACBOOK_PRO_GET_RASPBERRY_PI.ID,
			PromoType:          PROMO_MACBOOK_PRO_GET_RASPBERRY_PI.Type,
		},
	}
	ORDERS_GOOGLE_HOME_BY_2_GET_3 = dao.Orders{
		Model:         BASE_GORM_MODEL,
		UserID:        1,
		FinalAmount:   PRODUCT_GOOGLE_HOME.Price * 2,
		OrderProducts: ORDER_PRODUCTS_GOOGLE_HOME_BY_2_GET_3,
	}
	ORDER_PRODUCTS_GOOGLE_HOME_BY_2_GET_3 = []dao.OrderProducts{{
		Model:              BASE_GORM_MODEL,
		OrderID:            BASE_GORM_MODEL.ID,
		ProductID:          PRODUCT_GOOGLE_HOME.ID,
		Quantity:           3,
		Price:              PRODUCT_GOOGLE_HOME.Price,
		TotalPrice:         PRODUCT_GOOGLE_HOME.Price * 2,
		FreeGiftQuantity:   1,
		DiscountPercentage: 0,
		DiscountAmount:     0,
		FinalAmount:        PRODUCT_GOOGLE_HOME.Price * 2,
		IsPromo:            true,
		PromoID:            PROMO_GOOGLE_HOME_BUY_2_GET_3.ID,
		PromoType:          PROMO_GOOGLE_HOME_BUY_2_GET_3.Type,
	}}
	ORDERS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT_TOTAL_PRICE     = PRODUCT_ALEXA_SPEAKER.Price * 3
	ORDERS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT_DISCOUNT_AMOUNT = ORDERS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT_TOTAL_PRICE * (PROMO_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT.Gift.DiscountPercentage / 100)
	ORDERS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT_FINAL_AMOUNT    = ORDERS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT_TOTAL_PRICE - ORDERS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT_DISCOUNT_AMOUNT
	ORDERS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT                 = dao.Orders{
		Model:         BASE_GORM_MODEL,
		UserID:        1,
		FinalAmount:   ORDERS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT_FINAL_AMOUNT,
		OrderProducts: ORDER_PRODUCTS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT,
	}
	ORDER_PRODUCTS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT = []dao.OrderProducts{{
		Model:              BASE_GORM_MODEL,
		OrderID:            BASE_GORM_MODEL.ID,
		ProductID:          PRODUCT_ALEXA_SPEAKER.ID,
		Quantity:           3,
		Price:              PRODUCT_ALEXA_SPEAKER.Price,
		TotalPrice:         ORDERS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT_TOTAL_PRICE,
		FreeGiftQuantity:   0,
		DiscountPercentage: PROMO_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT.Gift.DiscountPercentage,
		DiscountAmount:     ORDERS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT_DISCOUNT_AMOUNT,
		FinalAmount:        ORDERS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT_FINAL_AMOUNT,
		IsPromo:            true,
		PromoID:            PROMO_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT.ID,
		PromoType:          PROMO_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT.Type,
	}}
)
