package repository

import (
	"context"
	"fmt"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/pkg/dto"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
)

func (rep *orderRepository) GetOrderList(ctx context.Context, dto *dto.PaginationDTO) []*models.Order {
	var orders []*models.Order

	if !rep.isValidSortColumn(dto.GetSortColumn()) {
		return orders
	}

	if !rep.isValidSortDirection(dto.GetSortType()) {
		return orders
	}

	rawQuery := `
select ag.id "agency.id", ag.name "agency.name", o.*
  from orders o
  join agencies ag on o.agency_id = ag.id 
                  and ag.deleted_at is null
 where o.deleted_at is null 
 order by o.%s %s
 limit $1 offset $2
`
	query := fmt.Sprintf(rawQuery, dto.GetSortColumn(), dto.GetSortType())

	err := rep.db.SelectContext(ctx, &orders, query, dto.GetLimitValue(), dto.GetOffsetValue())

	if err != nil {
		return []*models.Order{}
	}

	return orders
}

func (rep *orderRepository) isValidSortColumn(name string) bool {
	_, exists := map[string]struct{}{
		config.SortColumnID:         {},
		config.SortColumnStatus:     {},
		config.SortColumnRentalDate: {},
		config.SortColumnGuestCount: {},
		config.SortColumnCreatedAt:  {},
	}[name]

	return exists
}

func (rep *orderRepository) isValidSortDirection(name string) bool {
	_, exists := map[string]struct{}{
		config.SortDirectionASC:  {},
		config.SortDirectionDESC: {},
	}[name]

	return exists
}
