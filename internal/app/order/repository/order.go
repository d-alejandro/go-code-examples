package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
)

func (rep *orderRepository) GetOrder(ctx context.Context, id int) (*models.Order, error) {
	query := `
select ag.id "agency.id", ag.name "agency.name", o.*
  from orders o,
       agencies ag
 where o.agency_id = ag.id
   and o.id = $1
   and o.deleted_at is null
   and ag.deleted_at is null
 limit 1`

	var order models.Order

	err := rep.db.GetContext(ctx, &order, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	return &order, nil
}
