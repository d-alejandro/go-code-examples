package handler

import (
	"net/http"
	"strconv"

	"github.com/d-alejandro/go-code-examples/internal/app/order/presenter"
	"github.com/d-alejandro/go-code-examples/internal/app/order/request"
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/gin-gonic/gin"
)

type OrderUpdateHandler struct {
	useCase   OrderUpdateUseCaseInterface
	presenter OrderUpdatePresenterInterface
}

func NewOrderUpdateHandler(
	useCase OrderUpdateUseCaseInterface,
	presenterUpdate OrderUpdatePresenterInterface,
) *OrderUpdateHandler {
	return &OrderUpdateHandler{
		useCase:   useCase,
		presenter: presenterUpdate,
	}
}

func (handler *OrderUpdateHandler) Update(context *gin.Context) {
	paramID := context.Param("id")

	id, errParam := strconv.Atoi(paramID)
	if errParam != nil {
		presenter.PresentErrorPresenter(context, http.StatusBadRequest, config.MessageInvalidID)
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
