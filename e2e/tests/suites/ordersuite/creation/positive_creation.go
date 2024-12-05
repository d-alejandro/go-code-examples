package creation

import (
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (suite *OrderCreationSuite) TestPositiveCreation(t provider.T) {
	t.Title("Create order with positive test result")
	t.Parallel()

	var orderResponse *response.OrderResponse

	t.WithNewStep("Create order", func(stepCtx provider.StepCtx) {
		orderResponse = suite.Steps.CreateOrder(stepCtx)

		stepCtx.Require().NotNil(orderResponse)
		stepCtx.Require().NotEmpty(orderResponse.Data.ID)
	})

	t.WithNewStep("Show created order", func(stepCtx provider.StepCtx) {
		resp := suite.Steps.GetOrder(stepCtx, orderResponse.Data.ID)

		stepCtx.Require().NotNil(resp)
		stepCtx.Require().EqualValues(resp.Data, orderResponse.Data)
	})
}
