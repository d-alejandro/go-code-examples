package presenter

import (
	"github.com/d-alejandro/go-code-examples/internal/pkg/helpers"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

type OrderPresenter interface {
	PresentOrder(*gin.Context, *models.Order)
	PresentOrderList(*gin.Context, []*models.Order)

	PresentError(ctx *gin.Context, statusCode int, errors any)
}

type renderingHelper interface {
	RenderJSON(ctx *gin.Context, code int, obj any)
}

type orderPresenter struct {
	helpers.ErrorPresenter
	rendering renderingHelper
}

func NewOrderPresenter(rendering renderingHelper) OrderPresenter {
	presenter := &orderPresenter{
		rendering: rendering,
	}

	presenter.SetRenderingHelper(rendering)

	return presenter
}
