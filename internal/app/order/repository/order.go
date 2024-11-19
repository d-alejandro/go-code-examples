package repository

import "github.com/d-alejandro/go-code-examples/internal/pkg/models"

func (repository *orderRepository) GetOrder(id int) (*models.Order, error) {
	var order *models.Order

	conditions := `"` + models.TableOrders + `"."` + models.ColumnIDTableOrders + `" = ?`

	result := repository.gorm.
		Preload(models.RelationAgencyTableOrders).
		Take(&order, conditions, id)

	return order, result.Error
}
