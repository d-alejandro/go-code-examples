package api

import (
	"github.com/d-alejandro/go-code-examples/internal/app/http/requests"
	"github.com/d-alejandro/go-code-examples/internal/app/presenters"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderIndexController struct {
	useCase   OrderIndexUseCaseInterface
	presenter OrderIndexPresenterInterface
}

func NewOrderIndexController(
	useCase OrderIndexUseCaseInterface,
	presenter OrderIndexPresenterInterface,
) *OrderIndexController {
	return &OrderIndexController{
		useCase:   useCase,
		presenter: presenter,
	}
}

func (controller *OrderIndexController) Index(context *gin.Context) {
	var request requests.OrderIndexRequest

	err := request.Validate(context)
	if err != nil {
		presenters.PresentErrorPresenter(context, http.StatusBadRequest, err)
		return
	}

	response := controller.useCase.Execute(&request)

	controller.presenter.Present(context, response)
}
