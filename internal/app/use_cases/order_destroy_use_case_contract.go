package use_cases

import "github.com/d-alejandro/go-code-examples/internal/app/models"

type OrderShowDestroyUseCaseInterface interface {
	Execute(id int) (*models.Order, error)
}

type OrderDestroyRepositoryInterface interface {
	Make(*models.Order) error
}
