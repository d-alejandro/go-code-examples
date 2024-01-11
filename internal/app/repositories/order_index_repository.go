package repositories

import "gorm.io/gorm"

type OrderIndexRepository struct {
	database *gorm.DB
}

func NewOrderIndexRepository( /*database *gorm.DB*/ ) *OrderIndexRepository {
	return &OrderIndexRepository{
		//database: database,
	}
}

func (orderIndexRepository *OrderIndexRepository) Make() string {
	return "OrderIndexController run."
}
