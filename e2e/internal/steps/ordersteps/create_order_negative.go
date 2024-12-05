package ordersteps

import (
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/request"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (step *OrderSteps) CreateOrderNegative(stepCtx provider.StepCtx, req *request.OrderCreationRequest) error {
	orderResponse, err := step.client.Create(req)

	stepCtx.Require().Error(err)
	stepCtx.Require().Nil(orderResponse)

	return err
}
