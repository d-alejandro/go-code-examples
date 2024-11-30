package handler

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/d-alejandro/go-code-examples/internal/app/order/presenter"
	"github.com/d-alejandro/go-code-examples/internal/pkg/helpers"
	"github.com/d-alejandro/go-code-examples/internal/pkg/mocks"
	"github.com/d-alejandro/go-code-examples/internal/pkg/resource"
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

			require.Equal(t, order.ID, orderResource.ID)
		})

	present := presenter.NewOrderPresenter(mockRendering)
	validator := helpers.NewValidationHelper()

	handle := NewOrderHandler(uCase, present, validator)
	handle.Destroy(ctx)
}
