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

func TestPositiveDestroy(t *testing.T) {
	ctx := &gin.Context{}

	order, err := testhelpers.CreateOrder(ctx, repos)
	require.NoError(t, err)
	require.NotNil(t, order)

	ctx.AddParam("id", strconv.Itoa(order.ID))

	controller := gomock.NewController(t)
	defer controller.Finish()

	presenter := mocks.NewMockOrderPresenter(controller)

	presenter.
		EXPECT().
		PresentOrder(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ *gin.Context, order *models.Order) {
			require.NotNil(t, order)

			assert.Equal(t, order.ID, order.ID)
			assert.NotNil(t, order.DeletedAt)
		})

	validator := helpers.NewValidationHelper()

	handle := NewOrderHandler(uCase, presenter, validator)
	handle.Destroy(ctx)
}
