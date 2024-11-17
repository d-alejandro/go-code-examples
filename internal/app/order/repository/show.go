package repository

import (
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"gorm.io/gorm"
)

type OrderShowRepository struct {
	gorm *gorm.DB
}

func NewOrderShowRepository(grm *gorm.DB) *OrderShowRepository {
	return &OrderShowRepository{grm}
}

func (repository *OrderShowRepository) Make(id int) (*models.Order, error) {
	var order *models.Order

	conditions := `"` + models.TableOrders + `"."` + models.ColumnIDTableOrders + `" = ?`
	result := repository.gorm.
		Preload(models.RelationAgencyTableOrders).
		Take(&order, conditions, id)

	return order, result.Error
}
