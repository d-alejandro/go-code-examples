package api

import (
	"net/http"
	"strconv"

	"github.com/d-alejandro/go-code-examples/internal/app/http/requests"
	"github.com/d-alejandro/go-code-examples/internal/app/presenters"
	"github.com/gin-gonic/gin"
)

type OrderUpdateController struct {
	useCase   OrderUpdateUseCaseInterface
	presenter OrderUpdatePresenterInterface
}

func NewOrderUpdateController(
	useCase OrderUpdateUseCaseInterface,
	presenter OrderUpdatePresenterInterface,
) *OrderUpdateController {
	return &OrderUpdateController{
		useCase:   useCase,
		presenter: presenter,
	}
}

func (controller *OrderUpdateController) Update(context *gin.Context) {
	paramID := context.Param("id")

	id, errParam := strconv.Atoi(paramID)
	if errParam != nil {
		message := "The ID parameter is invalid."
		presenters.PresentErrorPresenter(context, http.StatusBadRequest, message)
		return
	}

	var request requests.OrderUpdateRequest

	err := request.Validate(context)
	if err != nil {
		presenters.PresentErrorPresenter(context, http.StatusBadRequest, err)
		return
	}

	response, useCaseError := controller.useCase.Execute(&request, id)
	if useCaseError != nil {
		presenters.PresentErrorPresenter(context, http.StatusBadRequest, useCaseError)
		return
	}

	controller.presenter.Present(context, response)
}
