package usecase

import "github.com/d-alejandro/go-code-examples/internal/app/order/models"

type OrderShowRepositoryInterface interface {
	Make(id int) (*models.Order, error)
}
