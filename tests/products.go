package tests

import (
	"time"

	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"
	"gorm.io/gorm"
)

var (
	PRODUCT_MACBOOK = dao.Products{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		SKU:      "MBP001",
		Name:     "MacBook Pro",
		Price:    35000000,
		Quantity: 10,
	}
	PRODUCT_RASPBERRY_PI = dao.Products{
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		SKU:      "RPI001",
		Name:     "Raspberry Pi B",
		Price:    1000000,
		Quantity: 10,
	}
	PRODUCT_GOOGLE_HOME = dao.Products{
		Model: gorm.Model{
			ID:        3,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		SKU:      "GOOGLE001",
		Name:     "Google Home",
		Price:    1000000,
		Quantity: 10,
	}
	PRODUCT_ALEXA_SPEAKER = dao.Products{
		Model: gorm.Model{
			ID:        4,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		SKU:      "ALEXA001",
		Name:     "Alexa Speaker",
		Price:    1000000,
		Quantity: 10,
	}
	PRODUCTS = []dao.Products{
		PRODUCT_MACBOOK,
		PRODUCT_RASPBERRY_PI,
		PRODUCT_GOOGLE_HOME,
		PRODUCT_ALEXA_SPEAKER,
	}
)
