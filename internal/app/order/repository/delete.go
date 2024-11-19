package repository

import "github.com/d-alejandro/go-code-examples/internal/pkg/models"

func (repository *orderRepository) Delete(order *models.Order) error {
	result := repository.gorm.Delete(order)
	return result.Error
}
