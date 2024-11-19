package bindings

import "github.com/d-alejandro/go-code-examples/internal/app/order/presenter"

type PresenterProvider struct {
	OrderPresenter presenter.OrderPresenter
}

func NewPresenterProvider() *PresenterProvider {
	return &PresenterProvider{
		OrderPresenter: presenter.NewOrderPresenter(),
	}
}
