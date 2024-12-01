package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/order/presenter"
	"github.com/d-alejandro/go-code-examples/internal/pkg/helpers"
)

type PresenterProvider struct {
	orderPresenter presenter.OrderPresenter
}

func NewPresenterProvider() *PresenterProvider {
	rendering := helpers.NewRenderingHelper()

	return &PresenterProvider{
		orderPresenter: presenter.NewOrderPresenter(rendering),
	}
}
