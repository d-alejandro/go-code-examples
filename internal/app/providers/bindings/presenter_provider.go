package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/helpers"
	"github.com/d-alejandro/go-code-examples/internal/app/presenters"
)

type PresenterProvider struct {
	OrderIndexPresenter *presenters.OrderIndexPresenter
	OrderPresenter      *presenters.OrderPresenter
}

func NewPresenterProvider(container *helpers.DependenciesContainer) *PresenterProvider {
	return &PresenterProvider{
		OrderIndexPresenter: presenters.NewOrderListPresenter(),
		OrderPresenter:      presenters.NewOrderPresenter(),
	}
}
