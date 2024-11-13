package api

import (
	"net/http"
	"strconv"

	"github.com/d-alejandro/go-code-examples/internal/app/presenters"
	"github.com/gin-gonic/gin"
)

type OrderDestroyController struct {
	useCase   OrderDestroyUseCaseInterface
	presenter OrderDestroyPresenterInterface
}

func NewOrderDestroyController(
	useCase OrderDestroyUseCaseInterface,
	presenter OrderDestroyPresenterInterface,
) *OrderDestroyController {
	return &OrderDestroyController{
		useCase:   useCase,
		presenter: presenter,
	}
}

func (controller *OrderDestroyController) Destroy(context *gin.Context) {
	paramID := context.Param("id")

	id, errParam := strconv.Atoi(paramID)
	if errParam != nil {
		message := "The ID parameter is invalid."
		presenters.PresentErrorPresenter(context, http.StatusBadRequest, message)
		return
	}

	response, useCaseError := controller.useCase.Execute(id)
	if useCaseError != nil {
		presenters.PresentErrorPresenter(context, http.StatusBadRequest, useCaseError)
		return
	}

	controller.presenter.Present(context, response)
}
