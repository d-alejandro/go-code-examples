package repository

import (
	"context"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/jmoiron/sqlx"
)

func (rep *orderRepository) Update(ctx context.Context, order *models.Order) error {
	query := `
update orders
   set guest_count = :guest_count,
       transport_count = :transport_count,
       user_name = :user_name,
       email = :email,
       phone = :phone,
       note = :note,
       admin_note = :admin_note,
       updated_at = :updated_at
 where id = :id
`
	namedQuery, args, err := sqlx.Named(query, order)

	if err != nil {
		return err
	}

	namedQuery = rep.db.Rebind(namedQuery)

	_, err = rep.db.ExecContext(ctx, namedQuery, args...)

	return err
}
