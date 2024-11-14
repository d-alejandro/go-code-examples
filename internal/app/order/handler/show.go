package handler

import (
	"net/http"
	"strconv"

	"github.com/d-alejandro/go-code-examples/internal/app/order/presenter"
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/gin-gonic/gin"
)

type OrderShowHandler struct {
	useCase   OrderShowUseCaseInterface
	presenter OrderShowPresenterInterface
}

func NewOrderShowHandler(
	useCase OrderShowUseCaseInterface,
	presenterShow OrderShowPresenterInterface,
) *OrderShowHandler {
	return &OrderShowHandler{
		useCase:   useCase,
		presenter: presenterShow,
	}
}

func (handler *OrderShowHandler) Show(context *gin.Context) {
	paramID := context.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		presenter.PresentErrorPresenter(context, http.StatusBadRequest, config.MessageInvalidID)
		return
	}

	response, useCaseError := handler.useCase.Execute(id)
	if useCaseError != nil {
		presenter.PresentErrorPresenter(context, http.StatusBadRequest, useCaseError)
		return
	}

	handler.presenter.Present(context, response)
}
