package tests

import (
	"time"

	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"
	"gorm.io/gorm"
)

var (
	PROMO_MACBOOK_PRO_GET_RASPBERRY_PI = dao.Promos{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Type: dao.PromoTypeFreeGift,
		Threshold: dao.PromoThreshold{
			ProductID:         PRODUCT_MACBOOK.ID,
			MinimumProductQty: 1,
		},
		Gift: dao.PromoGift{
			FreeGiftProductID:  PRODUCT_RASPBERRY_PI.ID,
			FreeGiftProductQty: 1,
		},
	}
	PROMO_GOOGLE_HOME_BUY_2_GET_3 = dao.Promos{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Type: dao.PromoTypeFixed,
		Threshold: dao.PromoThreshold{
			ProductID:         PRODUCT_GOOGLE_HOME.ID,
			MinimumProductQty: 2,
		},
		Gift: dao.PromoGift{
			FixedProductID:      PRODUCT_GOOGLE_HOME.ID,
			FixedGiftProductQty: 1,
		},
	}
	PROMO_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT = dao.Promos{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Type: dao.PromoTypeDiscount,
		Threshold: dao.PromoThreshold{
			ProductID:         PRODUCT_ALEXA_SPEAKER.ID,
			MinimumProductQty: 3,
		},
		Gift: dao.PromoGift{
			DiscountPercentage: 10,
		},
	}
	PROMOS = []dao.Promos{
		PROMO_MACBOOK_PRO_GET_RASPBERRY_PI,
		PROMO_GOOGLE_HOME_BUY_2_GET_3,
		PROMO_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT,
	}
	NO_PROMOS = []dao.Promos{}
)
