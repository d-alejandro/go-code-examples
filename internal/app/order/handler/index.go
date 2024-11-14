package handler

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/app/order/presenter"
	"github.com/d-alejandro/go-code-examples/internal/app/order/request"
	"github.com/gin-gonic/gin"
)

type OrderIndexHandler struct {
	useCase   OrderIndexUseCaseInterface
	presenter OrderIndexPresenterInterface
}

func NewOrderIndexHandler(
	useCase OrderIndexUseCaseInterface,
	indexPresenter OrderIndexPresenterInterface,
) *OrderIndexHandler {
	return &OrderIndexHandler{
		useCase:   useCase,
		presenter: indexPresenter,
	}
}

func (handler *OrderIndexHandler) Index(context *gin.Context) {
	var req request.OrderIndexRequest

	err := req.Validate(context)
	if err != nil {
		presenter.PresentErrorPresenter(context, http.StatusBadRequest, err)
		return
	}

	response := handler.useCase.Execute(&req)

	handler.presenter.Present(context, response)
}
