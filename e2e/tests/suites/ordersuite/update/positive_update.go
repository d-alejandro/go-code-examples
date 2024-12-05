package update

import (
	"github.com/d-alejandro/go-code-examples/e2e/internal/httpclient"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/request"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (suite *OrderUpdateSuite) TestPositiveUpdate(t provider.T) {
	t.Title("Update order with positive test result")
	t.Parallel()

	tt := httpclient.NewOrderClient()

	orderUpdateRequest := &request.OrderUpdateRequest{
		ID:             10000065,
		GuestCount:     1,
		TransportCount: 1,
		UserName:       nil,
		Email:          "test@test.ru",
		Phone:          "71437854000",
		Note:           nil,
		AdminNote:      nil,
	}

	resp, zz, err := tt.Update(orderUpdateRequest)

	if err != nil {
	}

	_ = zz
	_ = err
	_ = resp

}
