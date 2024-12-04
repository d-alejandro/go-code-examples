package ordersuite

import (
	"github.com/d-alejandro/go-code-examples/e2e/tests/suites/ordersuite/creation"
	"github.com/d-alejandro/go-code-examples/e2e/tests/suites/ordersuite/deletion"
	"github.com/d-alejandro/go-code-examples/e2e/tests/suites/ordersuite/reading"
	"github.com/d-alejandro/go-code-examples/e2e/tests/suites/ordersuite/update"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type OrderRunnerSuite struct {
	suite.Suite
}

func (suite *OrderRunnerSuite) TestOrderCreation(t provider.T) {
	t.Parallel()
	suite.RunSuite(t, new(creation.OrderCreationSuite))
}

func (suite *OrderRunnerSuite) TestOrderDeletion(t provider.T) {
	t.Parallel()
	suite.RunSuite(t, new(deletion.OrderDeletionSuite))
}

func (suite *OrderRunnerSuite) TestOrderReading(t provider.T) {
	t.Parallel()
	suite.RunSuite(t, new(reading.OrderReadingSuite))
}

func (suite *OrderRunnerSuite) TestOrderUpdate(t provider.T) {
	t.Parallel()
	suite.RunSuite(t, new(update.OrderUpdateSuite))
}
