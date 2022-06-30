package tests

import (
	"time"

	"gorm.io/gorm"
)

var (
	BASE_GORM_MODEL = gorm.Model{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
)
