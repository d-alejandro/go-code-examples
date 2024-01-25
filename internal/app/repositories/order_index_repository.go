package repositories

import (
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"gorm.io/gorm"
)

type OrderIndexRepository struct {
	gorm *gorm.DB
}

const (
	agencyRelation = "Agency"
)

func NewOrderIndexRepository(gorm *gorm.DB) *OrderIndexRepository {
	return &OrderIndexRepository{gorm}
}

func (repository *OrderIndexRepository) Make(pagination interface{ PaginationDTOInterface }) []models.Order {
	var orders []models.Order

	orderQueryRaw := pagination.GetSortColumn() + " " + pagination.GetSortType()

	repository.gorm.
		Preload(agencyRelation).
		Order(orderQueryRaw).
		Limit(pagination.GetLimitValue()).
		Offset(pagination.GetOffsetValue()).
		Find(&orders)

	return orders
}
