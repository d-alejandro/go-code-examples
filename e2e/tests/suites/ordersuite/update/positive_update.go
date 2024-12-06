package update

import (
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/request"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (suite *OrderUpdateSuite) TestPositiveUpdate(t provider.T) {
	t.Title("Update order with positive test result")
	t.Parallel()

	var orderResponse *response.OrderResponse

	t.WithNewStep("Create order", func(stepCtx provider.StepCtx) {
		orderResponse = suite.Steps.CreateOrder(stepCtx)

		stepCtx.Require().NotNil(orderResponse)
		stepCtx.Require().NotEmpty(orderResponse.Data.ID)
	})

	t.WithNewStep("Update order", func(stepCtx provider.StepCtx) {
		updateRequest := &request.OrderUpdateRequest{
			ID:             orderResponse.Data.ID,
			GuestCount:     orderResponse.Data.GuestCount + 2,
			TransportCount: orderResponse.Data.TransportCount + 1,
			UserName:       nil,
			Email:          orderResponse.Data.Email,
			Phone:          orderResponse.Data.Phone,
			Note:           nil,
			AdminNote:      nil,
		}

		updateResponse := suite.Steps.UpdateOrder(stepCtx, updateRequest)

		stepCtx.Require().NotNil(updateResponse)
	})

	t.WithNewStep("Show updated order", func(stepCtx provider.StepCtx) {
		resp := suite.Steps.GetOrder(stepCtx, orderResponse.Data.ID)

		stepCtx.Require().NotNil(resp)

		stepCtx.Assert().EqualValues(resp.Data.ID, orderResponse.Data.ID)
		stepCtx.Assert().EqualValues(resp.Data.TransportCount, orderResponse.Data.TransportCount+1)
		stepCtx.Assert().EqualValues(resp.Data.GuestCount, orderResponse.Data.GuestCount+2)
		stepCtx.Assert().Empty(resp.Data.UserName)
		stepCtx.Assert().EqualValues(resp.Data.Email, orderResponse.Data.Email)
		stepCtx.Assert().EqualValues(resp.Data.Phone, orderResponse.Data.Phone)
		stepCtx.Assert().Empty(resp.Data.Note)
		stepCtx.Assert().Empty(resp.Data.AdminNote)
	})
}
