package handler

import (
	"net/http"
	"strconv"

	"github.com/d-alejandro/go-code-examples/internal/app/order/presenter"
	"github.com/d-alejandro/go-code-examples/internal/app/order/request"
	"github.com/gin-gonic/gin"
)

type OrderUpdateHandler struct {
	useCase   OrderUpdateUseCaseInterface
	presenter OrderUpdatePresenterInterface
}

func NewOrderUpdateHandler(
	useCase OrderUpdateUseCaseInterface,
	presenter OrderUpdatePresenterInterface,
) *OrderUpdateHandler {
	return &OrderUpdateHandler{
		useCase:   useCase,
		presenter: presenter,
	}
}

func (handler *OrderUpdateHandler) Update(context *gin.Context) {
	paramID := context.Param("id")

	id, errParam := strconv.Atoi(paramID)
	if errParam != nil {
		message := "The ID parameter is invalid."
		presenter.PresentErrorPresenter(context, http.StatusBadRequest, message)
		return
	}

	var req request.OrderUpdateRequest

	err := req.Validate(context)
	if err != nil {
		presenter.PresentErrorPresenter(context, http.StatusBadRequest, err)
		return
	}

	response, useCaseError := handler.useCase.Execute(&req, id)
	if useCaseError != nil {
		presenter.PresentErrorPresenter(context, http.StatusBadRequest, useCaseError)
		return
	}

	handler.presenter.Present(context, response)
}
