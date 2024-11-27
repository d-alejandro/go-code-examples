package repository

import (
	"context"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
)

func (rep *orderRepository) Delete(ctx context.Context, order *models.Order) error {
	query := `
update orders
   set deleted_at = $1
 where id = $2
`
	_, err := rep.db.ExecContext(ctx, query, order.DeletedAt, order.ID)

	return err
}
