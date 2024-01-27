package repositories

import (
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderIndexRepository struct {
	gorm *gorm.DB
}

func NewOrderIndexRepository(gorm *gorm.DB) *OrderIndexRepository {
	return &OrderIndexRepository{gorm}
}

func (repository *OrderIndexRepository) Make(pagination interface{ PaginationDTOInterface }) []models.Order {
	var orders []models.Order

	column := clause.Column{
		Name: models.TableOrders + "." + pagination.GetSortColumn(),
	}

	const sortTypeDesc = "desc"
	orderByColumn := clause.OrderByColumn{
		Column: column,
		Desc:   pagination.GetSortType() == sortTypeDesc,
	}

	repository.gorm.
		Preload(models.RelationAgencyTableOrders).
		Order(orderByColumn).
		Limit(pagination.GetLimitValue()).
		Offset(pagination.GetOffsetValue()).
		Find(&orders)

	return orders
}
