package api

import (
	"github.com/d-alejandro/go-code-examples/internal/app/use_cases/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderIndexController struct {
	useCase interfaces.OrderIndexUseCaseInterface
}

func NewOrderIndexController(useCase interfaces.OrderIndexUseCaseInterface) *OrderIndexController {
	return &OrderIndexController{
		useCase: useCase,
	}
}

func (controller *OrderIndexController) Index(context *gin.Context) {
	context.String(http.StatusOK, controller.useCase.Execute())
}
