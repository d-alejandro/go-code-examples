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

func (runner *OrderRunnerSuite) TestOrderCreation(t provider.T) {
	t.Parallel()
	runner.RunSuite(t, new(creation.OrderCreationSuite))
}

func (runner *OrderRunnerSuite) TestOrderDeletion(t provider.T) {
	t.Parallel()
	runner.RunSuite(t, new(deletion.OrderDeletionSuite))
}

func (runner *OrderRunnerSuite) TestOrderReading(t provider.T) {
	t.Parallel()
	runner.RunSuite(t, new(reading.OrderReadingSuite))
}

func (runner *OrderRunnerSuite) TestOrderUpdate(t provider.T) {
	t.Parallel()
	runner.RunSuite(t, new(update.OrderUpdateSuite))
}
