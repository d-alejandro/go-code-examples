package handler

import (
	"github.com/jaswdr/faker/v2"
	"testing"
	"time"

	"github.com/d-alejandro/go-code-examples/internal/pkg/mocks"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestStore(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	fake := faker.New()

	mockValidator := mocks.NewMockValidationHelper(controller)

	mockValidator.
		EXPECT().
		ValidateForm(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx *gin.Context, req any) error {
			minDate := time.Now()
			maxDate := time.Now().AddDate(1, 0, 0)
			userName := fake.Person().Name()

			storeRequest := &request.OrderStoreRequest{
				AgencyName:     fake.Company().Name(),
				RentalDate:     fake.Time().TimeBetween(minDate, maxDate),
				GuestCount:     1,
				TransportCount: 1,
				UserName:       &userName,
				Email:          fake.Internet().Email(),
				Phone:          fake.Phone().Number(),
			}

			*(req.(*request.OrderStoreRequest)) = *storeRequest

			return nil
		})

	handle := NewOrderHandler(uCase, present, mockValidator)

	ctx := &gin.Context{}
	handle.Store(ctx)
}
