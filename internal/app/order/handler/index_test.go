package handler

import (
	"net/http"
	"testing"

	"github.com/d-alejandro/go-code-examples/internal/app/order/presenter"
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/pkg/mocks"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"github.com/d-alejandro/go-code-examples/internal/pkg/resource"
	"github.com/d-alejandro/go-code-examples/internal/pkg/testhelpers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestPositiveIndex(t *testing.T) {
	t.Run("Positive case", func(t *testing.T) {
		ctx := &gin.Context{}

		controller := gomock.NewController(t)
		defer controller.Finish()

		order, err := testhelpers.CreateOrder(ctx, repos)
		require.NoError(t, err)
		require.NotNil(t, order)

		orderIndexRequest := request.OrderIndexRequest{
			Start:      config.IndexNumberRequestStart,
			End:        config.IndexNumberRequestEnd,
			SortColumn: config.SortColumnID,
			SortType:   config.SortDirectionDESC,
			IDs:        []int{order.ID},
		}

		mockValidatorHelper := mocks.NewMockValidationHelper(controller)
		mockValidatorHelper.
			EXPECT().
			ValidateQuery(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ *gin.Context, req any) error {
				reqPointer, isOk := req.(*request.OrderIndexRequest)
				require.True(t, isOk)
				require.NotNil(t, reqPointer)

				*reqPointer = orderIndexRequest
				return nil
			})

		mockRenderingHelper := mocks.NewMockRenderingHelper(controller)
		mockRenderingHelper.
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

				orderResource, ok := resp.([]*resource.OrderIndexResource)
				require.True(t, ok)
				require.NotNil(t, orderResource)

				assert.Len(t, orderResource, 1)
				assert.Equal(t, order.ID, orderResource[0].ID)
				assert.Equal(t, string(models.Canceled), orderResource[0].Status)
			})

		present := presenter.NewOrderPresenter(mockRenderingHelper)

		handle := NewOrderHandler(uCase, present, mockValidatorHelper)
		handle.Index(ctx)
	})
}
