package repositories

import (
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderIndexRepository struct {
	gorm *gorm.DB
}

const (
	agencyRelation  = "Agency"
	tableNameOrders = "orders"
	sortTypeDesc    = "desc"
)

func NewOrderIndexRepository(gorm *gorm.DB) *OrderIndexRepository {
	return &OrderIndexRepository{gorm}
}

func (repository *OrderIndexRepository) Make(pagination interface{ PaginationDTOInterface }) []models.Order {
	var orders []models.Order

	column := clause.Column{
		Name: tableNameOrders + "." + pagination.GetSortColumn(),
	}

	orderByColumn := clause.OrderByColumn{
		Column: column,
		Desc:   pagination.GetSortType() == sortTypeDesc,
	}

	repository.gorm.
		Preload(agencyRelation).
		Order(orderByColumn).
		Limit(pagination.GetLimitValue()).
		Offset(pagination.GetOffsetValue()).
		Find(&orders)

	return orders
}
