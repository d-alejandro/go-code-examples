package resources

import "github.com/d-alejandro/go-code-examples/internal/app/models"

type OrderIndexResource struct {
	ID uint `json:"id"`
}

func NewOrderIndexResource(order models.Order) OrderIndexResource {
	return OrderIndexResource{
		ID: order.ID,
	}
}
