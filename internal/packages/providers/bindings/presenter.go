package bindings

import "github.com/d-alejandro/go-code-examples/internal/app/order/presenter"

type PresenterProvider struct {
	OrderIndexPresenter *presenter.OrderIndexPresenter
	OrderPresenter      *presenter.OrderPresenter
}

func NewPresenterProvider() *PresenterProvider {
	return &PresenterProvider{
		OrderIndexPresenter: presenter.NewOrderListPresenter(),
		OrderPresenter:      presenter.NewOrderPresenter(),
	}
}
