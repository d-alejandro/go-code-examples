package api

import (
	"github.com/d-alejandro/go-code-examples/internal/app/presenters"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	paramId := context.Param("id")

	id, errParam := strconv.Atoi(paramId)
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
