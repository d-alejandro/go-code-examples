package handler

import (
	"net/http"
	"testing"
	"time"

	"github.com/d-alejandro/go-code-examples/internal/app/order/presenter"
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/pkg/helpers"
	"github.com/d-alejandro/go-code-examples/internal/pkg/mocks"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"github.com/d-alejandro/go-code-examples/internal/pkg/resource"
	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestPositiveStore(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	fake := faker.New()

	minDate := time.Now()
	maxDate := minDate.AddDate(1, 0, 0)
	userName := fake.Person().Name()

	orderStoreRequest := request.OrderStoreRequest{
		AgencyName:     fake.Company().Name(),
		RentalDate:     fake.Time().TimeBetween(minDate, maxDate),
		GuestCount:     1,
		TransportCount: 1,
		UserName:       &userName,
		Email:          fake.Internet().Email(),
		Phone:          fake.Phone().Number(),
	}

	mockValidator := mocks.NewMockValidationHelper(controller)

	mockValidator.
		EXPECT().
		ValidateForm(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ *gin.Context, req any) error {
			reqPointer, isOk := req.(*request.OrderStoreRequest)
			require.True(t, isOk)
			require.NotNil(t, reqPointer)

			*reqPointer = orderStoreRequest
			return nil
		})

	mockRendering := mocks.NewMockRenderingHelper(controller)

	mockRendering.
		EXPECT().
		RenderJSON(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ *gin.Context, code int, obj any) {
			assert.Equal(t, http.StatusOK, code)

			responseBody, isOk := obj.(gin.H)
			require.True(t, isOk)
			require.NotNil(t, responseBody)

			resp, isExist := responseBody["data"]
			require.True(t, isExist)
			require.NotNil(t, resp)

			orderResource, ok := resp.(*resource.OrderResource)
			require.True(t, ok)
			require.NotNil(t, orderResource)

			assert.Equal(t, orderStoreRequest.AgencyName, orderResource.AgencyName)
			assert.Equal(
				t,
				helpers.ConvertDate[time.Time, string](orderStoreRequest.RentalDate, config.DateLayout),
				orderResource.RentalDate,
			)
			assert.Equal(t, orderStoreRequest.GuestCount, orderResource.GuestCount)
			assert.Equal(t, orderStoreRequest.TransportCount, orderResource.TransportCount)
			assert.Equal(t, orderStoreRequest.UserName, orderResource.UserName)
			assert.Equal(t, orderStoreRequest.Email, orderResource.Email)
			assert.Equal(t, orderStoreRequest.Phone, orderResource.Phone)
		})

	present := presenter.NewOrderPresenter(mockRendering)

	handle := NewOrderHandler(uCase, present, mockValidator)

	ctx := &gin.Context{}
	handle.Store(ctx)
}
