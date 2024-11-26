package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/d-alejandro/go-code-examples/internal/pkg/helpers"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/jmoiron/sqlx"
)

func (rep *orderRepository) Create(ctx context.Context, order *models.Order) error {
	return helpers.ExecuteWithinTransaction(ctx, rep.db, func(tx *sqlx.Tx) error {
		query := `
select id
  from agencies
 where name = $1
   and deleted_at is null
 limit 1`

		err := tx.GetContext(ctx, &order.Agency.ID, query, order.Agency.Name)

		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return err
			}

			if err = rep.createAgency(ctx, tx, &order.Agency); err != nil {
				return err
			}
		}

		order.AgencyID = order.Agency.ID

		return rep.createOrder(ctx, tx, order)
	})
}

func (rep *orderRepository) createAgency(ctx context.Context, tx *sqlx.Tx, agency *models.Agency) error {
	query := `
insert into agencies (name, created_at, updated_at)
values (:name, :created_at, :updated_at)
returning id
`
	namedQuery, args, err := sqlx.Named(query, agency)

	if err != nil {
		return err
	}

	namedQuery = tx.Rebind(namedQuery)

	return tx.GetContext(ctx, &agency.ID, namedQuery, args...)
}

func (rep *orderRepository) createOrder(ctx context.Context, tx *sqlx.Tx, order *models.Order) error {
	query := `
insert into orders (agency_id, status, rental_date, guest_count, transport_count, user_name, email, phone,
                    note, admin_note, created_at, updated_at)
values (:agency_id, :status, :rental_date, :guest_count, :transport_count, :user_name, :email, :phone,
        :note, :admin_note, :created_at, :updated_at)
returning id
`
	namedQuery, args, err := sqlx.Named(query, order)

	if err != nil {
		return err
	}

	namedQuery = tx.Rebind(namedQuery)

	return tx.GetContext(ctx, &order.ID, namedQuery, args...)
}
