package api

import (
	"github.com/d-alejandro/go-code-examples/internal/app/http/requests"
	"github.com/d-alejandro/go-code-examples/internal/app/presenters"
	"github.com/d-alejandro/go-code-examples/internal/app/use_cases/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	orderIndexRequest := new(requests.OrderIndexRequest)

	err := context.ShouldBindWith(orderIndexRequest, binding.Query)
	if err != nil {
		presenters.PresentErrorPresenter(context, http.StatusBadRequest, err)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"message": controller.useCase.Execute()},
	)
}
