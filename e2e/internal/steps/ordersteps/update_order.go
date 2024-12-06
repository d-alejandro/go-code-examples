package ordersteps

import (
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/request"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (step *OrderSteps) UpdateOrder(stepCtx provider.StepCtx, req *request.OrderUpdateRequest) *response.OrderResponse {
	orderResponse, err := step.client.Update(req)

	stepCtx.Require().NoError(err)
	stepCtx.Require().NotNil(orderResponse)

	return orderResponse
}
