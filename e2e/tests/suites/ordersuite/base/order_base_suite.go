package base

import (
	"github.com/d-alejandro/go-code-examples/e2e/internal/httpclients/orderclient"
	"github.com/d-alejandro/go-code-examples/e2e/internal/steps/ordersteps"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type OrderBaseSuite struct {
	suite.Suite

	Steps *ordersteps.OrderSteps
}

func (receiver *OrderBaseSuite) BeforeAll(provider.T) {
	client := orderclient.NewOrderClient()

	receiver.Steps = ordersteps.NewOrderSteps(client)
}
