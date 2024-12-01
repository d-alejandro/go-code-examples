package handler

import (
	"strconv"
	"testing"

	"github.com/d-alejandro/go-code-examples/internal/pkg/helpers"
	"github.com/d-alejandro/go-code-examples/internal/pkg/mocks"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/testhelpers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestPositiveShow(t *testing.T) {
	t.Run("Positive case", func(t *testing.T) {
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
			})

		validatorHelper := helpers.NewValidationHelper()

		handle := NewOrderHandler(uCase, present, validatorHelper)
		handle.Show(ctx)
	})
}
