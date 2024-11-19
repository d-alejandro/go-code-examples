package bindings

import "github.com/d-alejandro/go-code-examples/internal/app/order/presenter"

type PresenterProvider struct {
	orderPresenter presenter.OrderPresenter
}

func NewPresenterProvider() *PresenterProvider {
	return &PresenterProvider{
		orderPresenter: presenter.NewOrderPresenter(),
	}
}
