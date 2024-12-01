package repository

import (
	"context"
	"fmt"

	"github.com/d-alejandro/go-code-examples/internal/pkg/dto"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/jmoiron/sqlx"
)

func (rep *orderRepository) GetOrderList(ctx context.Context, paginationDTO *dto.PaginationDTO) []*models.Order {
	sortColumn := paginationDTO.GetSortColumn()
	if !rep.isValidSortColumn(sortColumn) {
		return nil
	}

	sortType := paginationDTO.GetSortType()
	if !rep.isValidSortDirection(sortType) {
		return nil
	}

	rawQuery := `
select ag.id "agency.id", ag.name "agency.name", ord.*
  from orders ord
  join agencies ag on ord.agency_id = ag.id
                  and ag.deleted_at is null
 where ord.deleted_at is null%s
 order by ord.%s %s
 limit :limit_value offset :offset_value
`
	var whereInIDsCriteria string

	if paginationDTO.GetIDs() != nil {
		whereInIDsCriteria = " and ord.id in (:ids)"
	}

	query := fmt.Sprintf(rawQuery, whereInIDsCriteria, sortColumn, sortType)

	namedQuery, args, err := sqlx.Named(query, paginationDTO)

	if err != nil {
		return nil
	}

	namedQuery, args, err = sqlx.In(namedQuery, args...)

	if err != nil {
		return nil
	}

	namedQuery = rep.db.Rebind(namedQuery)

	var orders []*models.Order

	err = rep.db.SelectContext(ctx, &orders, namedQuery, args...)

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
