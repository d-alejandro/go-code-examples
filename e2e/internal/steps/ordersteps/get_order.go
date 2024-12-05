package ordersteps

import (
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (step *OrderSteps) GetOrder(stepCtx provider.StepCtx, id int) (*response.OrderResponse, error) {
	orderResponse, err := step.client.Get(id)

	stepCtx.Require().NoError(err)
	stepCtx.Require().NotNil(orderResponse)

	return orderResponse, nil
}
