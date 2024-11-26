package repository

import (
	"context"
	"fmt"

	"github.com/d-alejandro/go-code-examples/internal/pkg/dto"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
)

func (rep *orderRepository) GetOrderList(ctx context.Context, paginationDTO *dto.PaginationDTO) []*models.Order {
	if !rep.isValidSortColumn(paginationDTO.GetSortColumn()) {
		return nil
	}

	if !rep.isValidSortDirection(paginationDTO.GetSortType()) {
		return nil
	}

	var orders []*models.Order

	rawQuery := `
select ag.id "agency.id", ag.name "agency.name", o.*
  from orders o
  join agencies ag on o.agency_id = ag.id
                  and ag.deleted_at is null
 where o.deleted_at is null
 order by o.%s %s
 limit $1 offset $2
`
	query := fmt.Sprintf(rawQuery, paginationDTO.GetSortColumn(), paginationDTO.GetSortType())

	err := rep.db.SelectContext(ctx, &orders, query, paginationDTO.GetLimitValue(), paginationDTO.GetOffsetValue())

	if err != nil {
		return nil
	}

	return orders
}

func (rep *orderRepository) isValidSortColumn(name string) bool {
	_, exists := rep.sortColumnMap[name]
	return exists
}

func (rep *orderRepository) isValidSortDirection(name string) bool {
	_, exists := rep.sortDirectionMap[name]
	return exists
}
