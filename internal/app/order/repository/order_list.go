package repository

import (
	"github.com/d-alejandro/go-code-examples/internal/pkg/dto"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"gorm.io/gorm/clause"
)

const sortTypeDesc = "desc"

func (repository *orderRepository) GetOrderList(pagination dto.PaginationDTO) []*models.Order {
	var orders []*models.Order

	column := clause.Column{
		Name: models.TableOrders + "." + pagination.GetSortColumn(),
	}

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
