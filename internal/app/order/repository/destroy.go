package repository

import (
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"gorm.io/gorm"
)

type OrderDestroyRepository struct {
	gorm *gorm.DB
}

func NewOrderDestroyRepository(grm *gorm.DB) *OrderDestroyRepository {
	return &OrderDestroyRepository{grm}
}

func (repository *OrderDestroyRepository) Make(order *models.Order) error {
	result := repository.gorm.Delete(order)
	return result.Error
}
