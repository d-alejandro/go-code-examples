package api

import (
	"net/http"
	"strconv"

	"github.com/d-alejandro/go-code-examples/internal/app/presenters"
	"github.com/gin-gonic/gin"
)

type OrderShowController struct {
	useCase   OrderShowUseCaseInterface
	presenter OrderShowPresenterInterface
}

func NewOrderShowController(
	useCase OrderShowUseCaseInterface,
	presenter OrderShowPresenterInterface,
) *OrderShowController {
	return &OrderShowController{
		useCase:   useCase,
		presenter: presenter,
	}
}

func (controller *OrderShowController) Show(context *gin.Context) {
	paramID := context.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
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
