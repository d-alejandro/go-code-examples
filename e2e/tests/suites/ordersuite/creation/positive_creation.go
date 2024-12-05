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

		stepCtx.Assert().EqualValues(resp.Data.ID, orderResponse.Data.ID)
		stepCtx.Assert().EqualValues(resp.Data.AgencyName, orderResponse.Data.AgencyName)
		stepCtx.Assert().EqualValues(resp.Data.Status, orderResponse.Data.Status)
		stepCtx.Assert().EqualValues(resp.Data.IsConfirmed, orderResponse.Data.IsConfirmed)
		stepCtx.Assert().EqualValues(resp.Data.IsChecked, orderResponse.Data.IsChecked)
		stepCtx.Assert().EqualValues(resp.Data.RentalDate, orderResponse.Data.RentalDate)
		stepCtx.Assert().EqualValues(*resp.Data.UserName, *orderResponse.Data.UserName)
		stepCtx.Assert().EqualValues(resp.Data.TransportCount, orderResponse.Data.TransportCount)
		stepCtx.Assert().EqualValues(resp.Data.GuestCount, orderResponse.Data.GuestCount)
		stepCtx.Assert().EqualValues(*resp.Data.AdminNote, *orderResponse.Data.AdminNote)
		stepCtx.Assert().EqualValues(*resp.Data.Note, *orderResponse.Data.Note)
		stepCtx.Assert().EqualValues(resp.Data.Email, orderResponse.Data.Email)
		stepCtx.Assert().EqualValues(resp.Data.Phone, orderResponse.Data.Phone)
	})
}
