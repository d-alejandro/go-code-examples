package repositories

import "gorm.io/gorm"

type OrderIndexRepository struct {
	gorm *gorm.DB
}

func NewOrderIndexRepository(gorm *gorm.DB) *OrderIndexRepository {
	return &OrderIndexRepository{
		gorm: gorm,
	}
}

func (repository *OrderIndexRepository) Make() string {
	return "OrderIndexController run."
}
