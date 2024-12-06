package reading

import (
	"github.com/d-alejandro/go-code-examples/e2e/internal/config"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/request"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (suite *OrderReadingSuite) TestPositivePaginateOrders(t provider.T) {
	t.Title("Paginate orders with positive test result")
	t.Parallel()

	var (
		orderResponseFirst  *response.OrderResponse
		orderResponseSecond *response.OrderResponse
	)

	t.WithNewStep("Create first order", func(stepCtx provider.StepCtx) {
		orderResponseFirst = suite.Steps.CreateOrder(stepCtx)

		stepCtx.Require().NotNil(orderResponseFirst)
		stepCtx.Require().NotEmpty(orderResponseFirst.Data.ID)
	})

	t.WithNewStep("Create second order", func(stepCtx provider.StepCtx) {
		orderResponseSecond = suite.Steps.CreateOrder(stepCtx)

		stepCtx.Require().NotNil(orderResponseSecond)
		stepCtx.Require().NotEmpty(orderResponseSecond.Data.ID)
	})

	t.WithNewStep("Paginate orders", func(stepCtx provider.StepCtx) {
		req := &request.OrderPaginationRequest{
			Start:      config.OrderStartPagination,
			End:        config.OrderEndPagination,
			SortColumn: config.SortColumnID,
			SortType:   config.SortDirectionDESC,
			IDs: []int{
				orderResponseFirst.Data.ID,
				orderResponseSecond.Data.ID,
			},
		}

		orderListResponse := suite.Steps.PaginateOrders(stepCtx, req)

		stepCtx.Require().NotNil(orderListResponse)
		stepCtx.Require().Len(orderListResponse.Data, config.NumberPagesIsTwo)
		stepCtx.Require().EqualValues(orderListResponse.Data[0].ID, orderResponseSecond.Data.ID)
		stepCtx.Require().EqualValues(orderListResponse.Data[1].ID, orderResponseFirst.Data.ID)
	})
}
