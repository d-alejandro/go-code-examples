package handler

import (
	"net/http"
	"strconv"

	"github.com/d-alejandro/go-code-examples/internal/app/order/presenter"
	"github.com/gin-gonic/gin"
)

type OrderShowHandler struct {
	useCase   OrderShowUseCaseInterface
	presenter OrderShowPresenterInterface
}

func NewOrderShowHandler(
	useCase OrderShowUseCaseInterface,
	presenter OrderShowPresenterInterface,
) *OrderShowHandler {
	return &OrderShowHandler{
		useCase:   useCase,
		presenter: presenter,
	}
}

func (handler *OrderShowHandler) Show(context *gin.Context) {
	paramID := context.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		message := "The ID parameter is invalid."
		presenter.PresentErrorPresenter(context, http.StatusBadRequest, message)
		return
	}

	response, useCaseError := handler.useCase.Execute(id)
	if useCaseError != nil {
		presenter.PresentErrorPresenter(context, http.StatusBadRequest, useCaseError)
		return
	}

	handler.presenter.Present(context, response)
}
