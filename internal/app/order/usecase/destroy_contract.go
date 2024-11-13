package usecase

import "github.com/d-alejandro/go-code-examples/internal/app/order/models"

type OrderShowDestroyUseCaseInterface interface {
	Execute(id int) (*models.Order, error)
}

type OrderDestroyRepositoryInterface interface {
	Make(*models.Order) error
}
