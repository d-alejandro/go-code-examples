package deletion

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/helpers"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (suite *OrderDeletionSuite) TestPositiveDeletion(t provider.T) {
	t.Title("Delete order with positive test result")
	t.Parallel()

	var orderResponse *response.OrderResponse

	t.WithNewStep("Create order", func(stepCtx provider.StepCtx) {
		orderResponse = suite.Steps.CreateOrder(stepCtx)

		stepCtx.Require().NotNil(orderResponse)
		stepCtx.Require().NotEmpty(orderResponse.Data.ID)
	})

	t.WithNewStep("Delete order", func(stepCtx provider.StepCtx) {
		orderResp := suite.Steps.DeleteOrder(stepCtx, orderResponse.Data.ID)

		stepCtx.Require().NotNil(orderResp)
		stepCtx.Require().NotEmpty(orderResp.Data.ID)
	})

	t.WithNewStep("Attempt to show a deleted order", func(stepCtx provider.StepCtx) {
		err := suite.Steps.GetOrderNegative(stepCtx, orderResponse.Data.ID)
		stepCtx.Require().Error(err)

		var httpErr *helpers.HTTPError

		stepCtx.Require().ErrorAs(err, &httpErr)
		stepCtx.Require().EqualValues(http.StatusBadRequest, httpErr.Code)
		stepCtx.Require().Contains(httpErr.Message, `order not found`)
	})
}
