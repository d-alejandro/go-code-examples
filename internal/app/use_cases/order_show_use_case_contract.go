package use_cases

import "github.com/d-alejandro/go-code-examples/internal/app/models"

type OrderShowRepositoryInterface interface {
	Make(id int) (*models.Order, error)
}
