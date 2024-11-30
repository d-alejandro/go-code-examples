package handler

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/d-alejandro/go-code-examples/internal/app/order/presenter"
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/pkg/helpers"
	"github.com/d-alejandro/go-code-examples/internal/pkg/mocks"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/testhelpers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestPositiveDestroy(t *testing.T) {
	ctx := &gin.Context{}

	order, err := testhelpers.CreateOrder(ctx, repos)
	require.NoError(t, err)
	require.NotNil(t, order)

	ctx.AddParam("id", strconv.Itoa(order.ID))

	controller := gomock.NewController(t)
	defer controller.Finish()

	present := mocks.NewMockOrderPresenter(controller)
	present.
		EXPECT().
		PresentOrder(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ *gin.Context, orderModel *models.Order) {
			require.NotNil(t, orderModel)

			assert.Equal(t, order.ID, orderModel.ID)
			assert.NotNil(t, orderModel.DeletedAt)
		})

	validatorHelper := helpers.NewValidationHelper()

	handle := NewOrderHandler(uCase, present, validatorHelper)
	handle.Destroy(ctx)
}

func TestNegativeDestroy(t *testing.T) {
	t.Run("Negative cases", func(t *testing.T) {
		t.Run("ID parameter is not set", func(t *testing.T) {
			ctx := &gin.Context{}

			controller := gomock.NewController(t)
			defer controller.Finish()

			mockRenderingHelper := mocks.NewMockRenderingHelper(controller)
			mockRenderingHelper.
				EXPECT().
				RenderJSON(gomock.Any(), gomock.Any(), gomock.Any()).
				DoAndReturn(func(_ *gin.Context, code int, obj any) {
					assert.Equal(t, http.StatusBadRequest, code)

					errorResponse, isOk := obj.(*helpers.ErrorResponse)
					require.True(t, isOk)
					require.NotNil(t, errorResponse)

					assert.Equal(t, http.StatusText(code), errorResponse.Message)
					assert.Equal(t, config.MessageInvalidID, errorResponse.Errors)
				})

			present := presenter.NewOrderPresenter(mockRenderingHelper)

			validatorHelper := helpers.NewValidationHelper()

			handle := NewOrderHandler(uCase, present, validatorHelper)
			handle.Destroy(ctx)
		})
	})
}
