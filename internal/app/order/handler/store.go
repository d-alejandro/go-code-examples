package handler

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/app/order/presenter"
	"github.com/d-alejandro/go-code-examples/internal/app/order/request"
	"github.com/gin-gonic/gin"
)

type OrderStoreHandler struct {
	useCase   OrderStoreUseCaseInterface
	presenter OrderStorePresenterInterface
}

func NewOrderStoreHandler(
	useCase OrderStoreUseCaseInterface,
	presenter OrderStorePresenterInterface,
) *OrderStoreHandler {
	return &OrderStoreHandler{
		useCase:   useCase,
		presenter: presenter,
	}
}

func (handler *OrderStoreHandler) Store(context *gin.Context) {
	var req request.OrderStoreRequest

	err := req.Validate(context)
	if err != nil {
		presenter.PresentErrorPresenter(context, http.StatusBadRequest, err)
		return
	}

	response, useCaseError := handler.useCase.Execute(&req)
	if useCaseError != nil {
		presenter.PresentErrorPresenter(context, http.StatusBadRequest, useCaseError)
		return
	}

	handler.presenter.Present(context, response)
}
