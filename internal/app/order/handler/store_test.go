package handler

import (
	"testing"
	"time"

	"github.com/d-alejandro/go-code-examples/internal/app/order/presenter"
	"github.com/d-alejandro/go-code-examples/internal/pkg/mocks"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestStore(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockValidator := mocks.NewMockValidationHelper(controller)

	mockValidator.
		EXPECT().
		ValidateForm(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ *gin.Context, req any) error {
			reqPointer, isOk := req.(*request.OrderStoreRequest)
			require.True(t, isOk)

			fake := faker.New()

			minDate := time.Now()
			maxDate := time.Now().AddDate(1, 0, 0)
			userName := fake.Person().Name()

			*reqPointer = request.OrderStoreRequest{
				AgencyName:     fake.Company().Name(),
				RentalDate:     fake.Time().TimeBetween(minDate, maxDate),
				GuestCount:     1,
				TransportCount: 1,
				UserName:       &userName,
				Email:          fake.Internet().Email(),
				Phone:          fake.Phone().Number(),
			}

			return nil
		})

	mockRendering := mocks.NewMockRenderingHelper(controller)

	mockRendering.
		EXPECT().
		RenderJSON(gomock.Any(), gomock.Any(), gomock.Any()).
		Return()

	present := presenter.NewOrderPresenter(mockRendering)

	handle := NewOrderHandler(uCase, present, mockValidator)

	ctx := &gin.Context{}
	handle.Store(ctx)
}
