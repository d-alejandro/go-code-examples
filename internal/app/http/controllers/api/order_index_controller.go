package api

import (
	"github.com/d-alejandro/go-code-examples/internal/app/http/requests"
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"github.com/d-alejandro/go-code-examples/internal/app/presenters"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderIndexUseCase interface {
	Execute(request *requests.OrderIndexRequest) []models.Order
}

type OrderIndexController struct {
	useCase OrderIndexUseCase
}

func NewOrderIndexController(useCase OrderIndexUseCase) *OrderIndexController {
	return &OrderIndexController{useCase}
}

func (controller *OrderIndexController) Index(context *gin.Context) {
	var request requests.OrderIndexRequest

	err := request.Validate(context)
	if err != nil {
		presenters.PresentErrorPresenter(context, http.StatusBadRequest, err)
		return
	}

	response := controller.useCase.Execute(&request)

	context.JSON(
		http.StatusOK,
		gin.H{"data": response},
	)
}
