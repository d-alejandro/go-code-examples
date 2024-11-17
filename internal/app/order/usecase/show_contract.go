package usecase

import "github.com/d-alejandro/go-code-examples/internal/pkg/models"

type OrderShowRepositoryInterface interface {
	Make(id int) (*models.Order, error)
}
