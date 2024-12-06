package ordersteps

import "github.com/ozontech/allure-go/pkg/framework/provider"

func (step *OrderSteps) GetOrderNegative(stepCtx provider.StepCtx, id int) error {
	orderResponse, err := step.client.Get(id)

	stepCtx.Require().Error(err)
	stepCtx.Require().Nil(orderResponse)

	return err
}
