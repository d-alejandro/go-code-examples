package creation

import (
	"net/http"
	"time"

	"github.com/d-alejandro/go-code-examples/e2e/internal/config"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/helpers"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/request"
	"github.com/jaswdr/faker/v2"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (suite *OrderCreationSuite) TestNegativeCreation(t provider.T) {
	t.Title("Create order with negative test result")
	t.Parallel()

	t.WithNewStep(`Validation error: 'RentalDate' failed`, func(stepCtx provider.StepCtx) {
		fake := faker.New()

		req := &request.OrderCreationRequest{
			AgencyName:     fake.Company().Name(),
			RentalDate:     time.Time{},
			GuestCount:     config.GuestCountMin,
			TransportCount: config.TransportCountMin,
			Email:          fake.Internet().Email(),
			Phone:          fake.Phone().Number(),
		}

		err := suite.Steps.CreateOrderNegative(stepCtx, req)
		stepCtx.Require().Error(err)

		var httpErr *helpers.HTTPError

		stepCtx.Require().ErrorAs(err, &httpErr)
		stepCtx.Require().EqualValues(http.StatusBadRequest, httpErr.Code)
		stepCtx.Require().Contains(httpErr.Message, `Field validation for 'RentalDate' failed`)
	})
}
