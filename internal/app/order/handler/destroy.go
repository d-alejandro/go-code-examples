package handler

import (
	"net/http"
	"strconv"

	"github.com/d-alejandro/go-code-examples/internal/app/order/presenter"
	"github.com/gin-gonic/gin"
)

type OrderDestroyHandler struct {
	useCase   OrderDestroyUseCaseInterface
	presenter OrderDestroyPresenterInterface
}

func NewOrderDestroyHandler(
	useCase OrderDestroyUseCaseInterface,
	presenter OrderDestroyPresenterInterface,
) *OrderDestroyHandler {
	return &OrderDestroyHandler{
		useCase:   useCase,
		presenter: presenter,
	}
}

func (handler *OrderDestroyHandler) Destroy(context *gin.Context) {
	paramID := context.Param("id")

	id, errParam := strconv.Atoi(paramID)
	if errParam != nil {
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
