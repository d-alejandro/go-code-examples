package handler

import (
	"strconv"
	"testing"
	"time"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/pkg/mocks"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"github.com/d-alejandro/go-code-examples/internal/pkg/testhelpers"
	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestPositiveUpdate(t *testing.T) {
	t.Run("Positive case", func(t *testing.T) {
		ctx := &gin.Context{}

		order, err := testhelpers.CreateOrder(ctx, repos)
		require.NoError(t, err)
		require.NotNil(t, order)

		ctx.AddParam("id", strconv.Itoa(order.ID))

		controller := gomock.NewController(t)
		defer controller.Finish()

		fake := faker.New()

		userName := fake.Person().Name()

		orderUpdateRequest := request.OrderUpdateRequest{
			GuestCount:     config.GuestCountMin + 1,
			TransportCount: config.TransportCountMin + 1,
			UserName:       &userName,
			Email:          order.Email,
			Phone:          order.Phone,
		}

		mockValidatorHelper := mocks.NewMockValidationHelper(controller)
		mockValidatorHelper.
			EXPECT().
			ValidateForm(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ *gin.Context, req any) error {
				reqPointer, isOk := req.(*request.OrderUpdateRequest)
				require.True(t, isOk)
				require.NotNil(t, reqPointer)

				*reqPointer = orderUpdateRequest
				return nil
			})

		present := mocks.NewMockOrderPresenter(controller)
		present.
			EXPECT().
			PresentOrder(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ *gin.Context, orderModel *models.Order) {
				require.NotNil(t, orderModel)

				assert.Equal(t, order.ID, orderModel.ID)
				assert.Equal(t, orderUpdateRequest.GuestCount, orderModel.GuestCount)
				assert.Equal(t, orderUpdateRequest.TransportCount, orderModel.TransportCount)
				assert.Equal(t, orderUpdateRequest.UserName, orderModel.UserName)
				assert.Equal(t, orderUpdateRequest.Email, orderModel.Email)
				assert.Equal(t, orderUpdateRequest.Phone, orderModel.Phone)

				orderModel.CreatedAt = time.Date(
					orderModel.CreatedAt.Year(),
					orderModel.CreatedAt.Month(),
					orderModel.CreatedAt.Day(),
					orderModel.CreatedAt.Hour(),
					orderModel.CreatedAt.Minute(),
					orderModel.CreatedAt.Second(),
					0,
					time.Local,
				)

				assert.Equal(t, order.CreatedAt.Round(time.Second), orderModel.CreatedAt)
				assert.Nil(t, orderModel.DeletedAt)
			})

		handle := NewOrderHandler(uCase, present, mockValidatorHelper)
		handle.Update(ctx)
	})
}
