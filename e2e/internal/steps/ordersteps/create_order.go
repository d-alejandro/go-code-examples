package ordersteps

import (
	"time"

	"github.com/d-alejandro/go-code-examples/e2e/internal/config"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/request"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
	"github.com/jaswdr/faker/v2"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (step *OrderSteps) CreateOrder(stepCtx provider.StepCtx) *response.OrderResponse {
	fake := faker.New()

	minDate := time.Now()
	maxDate := minDate.AddDate(config.DateYearsMin, config.DateMonthsMin, config.DateDaysMin)
	userName := fake.Person().Name()

	note := fake.Letter()
	adminNote := fake.Letter()

	orderCreationRequest := &request.OrderCreationRequest{
		AgencyName:     fake.Company().Name(),
		RentalDate:     fake.Time().TimeBetween(minDate, maxDate),
		GuestCount:     config.GuestCountMin,
		TransportCount: config.TransportCountMin,
		UserName:       &userName,
		Email:          fake.Internet().Email(),
		Phone:          fake.Phone().Number(),
		Note:           &note,
		AdminNote:      &adminNote,
	}

	orderResponse, err := step.client.Create(orderCreationRequest)

	stepCtx.Require().NoError(err)
	stepCtx.Require().NotNil(orderResponse)

	return orderResponse
}
