package ordersteps

import (
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/request"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
)

type OrderSteps struct {
	client orderClient
}

type orderClient interface {
	Create(*request.OrderCreationRequest) (*response.OrderResponse, error)
	Delete(id int) (*response.OrderResponse, error)
	Get(id int) (*response.OrderResponse, error)
	Paginate(*request.OrderPaginationRequest) (*response.OrderListResponse, error)
	Update(*request.OrderUpdateRequest) (*response.OrderResponse, error)
}

func NewOrderSteps(client orderClient) *OrderSteps {
	return &OrderSteps{
		client: client,
	}
}
